from os import path
from google.genai import types
import subprocess

schema_run_python_file = types.FunctionDeclaration(
    name="run_python_file",
    description="Executes a Python file.",
    parameters=types.Schema(
        type=types.Type.OBJECT,
        properties={
            "file_path": types.Schema(
                type=types.Type.STRING,
                description="Path to the file to be executed.",
            ),
            "args": types.Schema(
                type=types.Type.ARRAY,
                items=types.Schema(type=types.Type.STRING),
                description="List of arguments to pass to the Python interpreter.",
            ),
        },
    ),
)

def run_python_file(working_directory, file_path, args=[]):
    abs_path = path.abspath(path.join(working_directory, file_path))
    abs_working = path.abspath(working_directory)

    if not abs_path.startswith(abs_working):
        return f'Error: Cannot execute "{file_path}" as it is outside the permitted working directory'

    if not path.isfile(abs_path):
        return f'Error: File "{file_path}" not found'

    if not abs_path.endswith(".py"):
        return f'Error: "{file_path}" is not a Python file'

    full_args = ["python", abs_path] + args 

    try:
        subp = subprocess.run(full_args, timeout=30, text=True, capture_output=True, cwd=abs_working)
        std_string = f'STDOUT: {subp.stdout}\nSTDERR: {subp.stderr}'

        if subp.returncode != 0:
            std_string += f'Process exited with code {subp.returncode}'

        return std_string

    except Exception as e:
        return f'Error executing Python file: {e}'
