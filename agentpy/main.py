import os
import sys

from dotenv import load_dotenv
from google import genai
from google.genai import types

from call_function import call_function, available_functions
from constants import system_prompt, model_name, MAX_ITERS

def main():
    load_dotenv()

    verbose = "--verbose" in sys.argv
    args = []
    for arg in sys.argv[1:]:
        if not arg.startswith("--"):
            args.append(arg)

    if not args:
        print("AI Code Assistant")
        print('\nUsage: python main.py "your prompt here" [--verbose]')
        sys.exit(1)

    api_key = os.environ.get("GEMINI_API_KEY")
    client = genai.Client(api_key=api_key)

    prompt = " ".join(args)
    if verbose:
        print(f"User prompt: {prompt}\n")

    history = [
        types.Content(role="user", parts=[types.Part(text=prompt)]),
    ]

    iters = 0
    while True:
        iters += 1
        if iters > MAX_ITERS:
            print(f"Maximum iterations ({MAX_ITERS}) reached.")
            sys.exit(1)
        try:
            response = generate_content(client, history, verbose)
            if response:
                print("Final Response:")
                print(response)
                break
        except Exception as e:
            print(f"Error in generate_content: {e}")

def generate_content(client, history, verbose):
    response = client.models.generate_content(
        model=model_name, 
        contents=history,
        config=genai.types.GenerateContentConfig(
            tools=[available_functions], 
            system_instruction=system_prompt
        ),
    )
    
    if verbose:
        print(f'Prompt tokens: {response.usage_metadata.prompt_token_count}')
        print(f'Response tokens: {response.usage_metadata.candidates_token_count}')
    
    if response.candidates:
        for candidate in response.candidates:
            history.append(candidate.content)

    if not response.function_calls:
        return response.text

    func_history = []
    for func_call in response.function_calls:
        func_response = call_function(func_call, verbose)
        if not func_response.parts[0].function_response.response:
            raise Exception("Empty function call result")
        if verbose:
            print(f"-> {func_response.parts[0].function_response.response}")
        func_history.append(func_response.parts[0])

    if not func_history:
        raise Exception("No function responses generated, exiting.")

    history.append(types.Content(role="user", parts=func_history))

if __name__ == "__main__":
    main()
