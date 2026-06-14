from os import path, SEEK_END 
from google.genai import types

MAX_CHARS = 10_000

schema_get_file_content = types.FunctionDeclaration(
    name="get_file_content",
    description="Returns the content of a file given its path.",
    parameters=types.Schema(
        type=types.Type.OBJECT,
        properties={
            "file_path": types.Schema(
                type=types.Type.STRING,
                description="Path to the file to get contents from",
            ),
        },
    ),
)

def get_file_content(working_directory, file_path):
    abs_path = path.abspath(path.join(working_directory, file_path))
    abs_working = path.abspath(working_directory)

    if not abs_path.startswith(abs_working):
        return f'Error: Cannot read "{file_path}" as it is outside the permitted working directory'

    if not path.isfile(abs_path):
        return f'Error: File not found or is not a regular file: "{file_path}"'

    try:
        with open(abs_path, "r") as f:
            content = f.read(MAX_CHARS) 
            if path.getsize(abs_path) > MAX_CHARS:
                content += f'[...File "{file_path}" truncated at 10000 characters]'
        return content
    except Exception as e:
        return f'Error reading file "{file_path}": {e}'
