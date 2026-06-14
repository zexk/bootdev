from stats import num_words
from stats import num_char
from stats import sort_list
import sys

def get_book_text(filepath):
    buffer = ""
    with open(filepath) as f:
        buffer = f.read()
    return buffer

def main():
    if len(sys.argv) != 2:
        print("Usage: python3 main.py <path_to_book>")
        sys.exit(1)

    buffer = get_book_text(sys.argv[1]) 
    char_count = sort_list(num_char(buffer))

    print("============ BOOKBOT ============")
    print(f"Analyzing book found at {sys.argv[1]}")

    print("----------- Word Count ----------")
    print(f"Found {num_words(buffer)} total words")

    print("--------- Character Count -------")
    for dict in char_count:
        char = dict["char"]
        count = dict["num"]
        print(f"{char}: {count}")

    print("============= END ===============")

main()
