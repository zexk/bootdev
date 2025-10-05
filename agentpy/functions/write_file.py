from os import path, makedirs 
from google.genai import types

schema_write_file = types.FunctionDeclaration(
    name="write_file",
    description="Writes a file containing given content.",
    parameters=types.Schema(
        type=types.Type.OBJECT,
        properties={
            "file_path": types.Schema(
                type=types.Type.STRING,
                description="Path to the file",
            ),
            "content": types.Schema(
                type=types.Type.STRING,
                description="Content to write in file",
            ),
        },
    ),
)

def write_file(working_directory, file_path, content):
    abs_file_path = path.abspath(path.join(working_directory, file_path))
    abs_working_dir = path.abspath(working_directory)

    if not abs_file_path.startswith(abs_working_dir):
        return f'Error: Cannot write to "{file_path}" as it is outside the permitted working directory'

    if not path.exists(abs_file_path):
        try:
            makedirs(path.dirname(abs_file_path), exist_ok=True)
        except Exception as e:
            return f"Error creating directory {e}"

    if path.exists(abs_file_path) and path.isdir(abs_file_path):
        return f'Error: "{file_path}" is a directory, not a file'

    try:
        with open(abs_file_path , "w") as f:
            f.write(content)
        return f'Successfully wrote to "{file_path}" ({len(content)} characters written)'
    except Exception as e:
        return f"Error: writing to file: {e}"
