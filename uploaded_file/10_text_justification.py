def text_justify(words, max_width):
    lines = []
    curr_line = []
    one_line_letters = 0
    for word in words:
        if len(word) + one_line_letters + len(curr_line) > max_width:
            for i in range(len(curr_line)):
                curr_line[i % (len(curr_line) - 1 or 1)] += ' '
            lines.append(''.join(curr_line))
            curr_line = []
            one_line_letters = 0
        curr_line.append(word)
        one_line_letters += len(word)
    
    return lines.append(' '.join(curr_line).ljust(max_width))