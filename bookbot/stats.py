def num_words(buffer):
    return len(buffer.split())    

def num_char(buffer):
    char_dict = {}
    buffer = buffer.lower()
    for i in range(len(buffer)):
        if buffer[i] not in char_dict:
            char_dict[buffer[i]] = 1
        else:
            char_dict[buffer[i]] += 1
    return char_dict

def sort_on(items):
    return items["num"]

def sort_list(char_dict):
    list = []
    for char in char_dict:
        if not char.isalpha():
            continue
        temp_dict = {}
        temp_dict["char"] = char 
        temp_dict["num"] = char_dict[char]
        list.append(temp_dict)
    list.sort(reverse=True, key=sort_on)
    return list
    
