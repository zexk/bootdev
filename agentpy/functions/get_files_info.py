from os import path, listdir
from google.genai import types

schema_get_files_info = types.FunctionDeclaration(
    name="get_files_info",
    description="Lists files in the specified directory along with their sizes, constrained to the working directory.",
    parameters=types.Schema(
        type=types.Type.OBJECT,
        properties={
            "directory": types.Schema(
                type=types.Type.STRING,
                description="The directory to list files from, relative to the working directory. If not provided, lists files in the working directory itself.",
            ),
        },
    ),
)

def get_files_info(working_directory, directory="."):
    abs_path = path.abspath(path.join(working_directory, directory))
    abs_working = path.abspath(working_directory)

    if not abs_path.startswith(abs_working):
        return f'Error: Cannot list "{directory}" as it is outside the permitted working directory'

    if not path.isdir(abs_path):
        return f'Error: "{directory}" is not a directory'

    try: 
        file_info = []
        for file in listdir(abs_path):
            file_path = path.join(abs_path, file)
            string = f'- {file}: file_size={path.getsize(file_path)} bytes, is_dir={path.isdir(file_path)}'    
            file_info.append(string)
        return "\n".join(file_info)
    except Exception as e:
        return f"Error listing files: {e}"
