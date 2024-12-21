from collections import Counter

from hmac import new
import sys
import os

f = open("input.txt", "r")
data = f.read().strip()
lines = data.split("\n")

def get_move(dir):
    if dir == "^":
        return -1, 0
    if dir == 'v':
        return 1, 0
    if dir == '<':
        return 0, -1
    if dir == '>':
        return 0, 1

def print_board(board):
    o = ''
    for i in range(len(board)):
        for j in range(len(board[0])):
            o += board[i][j]
        o += '\n'
    print(o)

board = []
moves = ''
for line in lines:
    if line == '':
        continue
    elif line[0] == '#':
        board.append(list(line))
    else:
        moves += line


def part1(board):
    L = len(board)
    C = len(board[0])
    print(board)
    
    sp = None
    for i in range(L):
        for j in range(C):
            if board[i][j] == '@':
                sp = (i, j)
                break

    print(sp)
    current = sp
    for move in moves:
        # print_board(board)
        print("next move is")
        print(move)
        x,y = get_move(move)
        next = (current[0] + x, current[1] + y)
        if board[next[0]][next[1]] == '#':
            continue
        elif board[next[0]][next[1]] == '.':
            board[current[0]][current[1]] = '.'
            current = next
            board[current[0]][current[1]] = '@'
            continue
        elif board[next[0]][next[1]] == 'O':
            allos = []
            while board[next[0]][next[1]] == 'O':
                allos.append(next)
                next = (next[0] + x, next[1] + y)
            print(allos, next, current)
            if board[next[0]][next[1]] == '#':
                continue
            elif board[next[0]][next[1]] == '.':
                board[current[0]][current[1]] = '.'
                first = allos[0]
                board[next[0]][next[1]] = 'O'
                board[first[0]][first[1]] = '@'
                current = first
    print(current)

    t = 0
    for i in range(L):
        for j in range(C):
            if board[i][j] == 'O':
                t += 100 * i +j
    print(t)

def part2(board):
    newb = []
    for i in range(len(board)):
        row = []
        for j in range(len(board[0])):
            if board[i][j] == '@':
                row.append('@')
                row.append('.')
            elif board[i][j] == 'O':
                row.append('[')
                row.append(']')
            else:
                row.append(board[i][j])
                row.append(board[i][j])
        newb.append(row)
    board = newb
    print('initial board')
    print_board(board)

    L = len(board)
    C = len(board[0])
 
    sp = None
    for i in range(L):
        for j in range(C):
            if board[i][j] == '@':
                sp = (i, j)
                break

    print(sp)
    current = sp
    for move in moves:
        # print_board(board)
        print("next move is")
        print(move)
        x,y = get_move(move)
        next = (current[0] + x, current[1] + y)

        if board[next[0]][next[1]] == '#':
            continue

        elif board[next[0]][next[1]] == '.':
            board[current[0]][current[1]] = '.'
            current = next
            board[current[0]][current[1]] = '@'
            continue

        # start implementing from here
        elif (move == '<' or move == '>') and (board[next[0]][next[1]] == '[' or board[next[0]][next[1]] == ']'):
            tomove = []
            while board[next[0]][next[1]] in '[]':
                tomove.append((next, board[next[0]][next[1]]))
                next = (next[0] + x, next[1] + y)
            # print(tomove, next, current)
            
            moved = []
            canmove = True
            for tm in tomove:
                pos, val = tm
                npos = (pos[0] + x, pos[1] + y)
                if board[npos[0]][npos[1]] == '#':
                    canmove = False
                    break
                moved.append((npos, val))

            if not canmove:
                continue
            
            # move everything to its new position, after removing them from
            # the old position. At the end, move the first one too
            for tm in tomove:
                pos, _ = tm
                board[pos[0]][pos[1]] = '.'
            
            for mo in moved:
                pos, val = mo
                board[pos[0]][pos[1]] = val

            board[current[0]][current[1]] = '.'
            board[current[0] + x][current[1] + y] = '@'

            current = (current[0]+x, current[1]+y)

        elif (move == '^' or move == 'v') and (board[next[0]][next[1]] == '[' or board[next[0]][next[1]] == ']'):
            
            tomove = [(current, '@')]
            if move == '^':
                for i in range(current[0]-1, -1, -1):
                    for j in range(0, C):
                        # those can lift up
                        tobelifted = [tm[0] for tm in tomove]
                        # check if they sit on something that is to be lifted
                        if board[i][j] == '[' and (((i+1,j) in tobelifted) or ((i+1,j + 1) in tobelifted)):
                            tomove.append(((i, j), '['))
                            tomove.append(((i, j+1), ']'))
            elif move == 'v':
                for i in range(current[0]+1, L, 1):
                    for j in range(0, C):
                        # those can lift down
                        tobelifted = [tm[0] for tm in tomove]
                        # check if they sit on something that is to be lifted
                        if board[i][j] == '[' and (((i-1,j) in tobelifted) or ((i-1,j + 1) in tobelifted)):
                            tomove.append(((i, j), '['))
                            tomove.append(((i, j+1), ']'))

            moved = []
            canmove = True
            for tm in tomove:
                pos, val = tm
                npos = (pos[0] + x, pos[1] + y)
                if board[npos[0]][npos[1]] == '#':
                    canmove = False
                    break
                moved.append((npos, val))

            if not canmove:
                continue
            
            # move everything to its new position, after removing them from
            # the old position. At the end, move the first one too.
            # This is identical from the first part, building the list is the
            # only thing that is different.
            for tm in tomove:
                pos, _ = tm
                board[pos[0]][pos[1]] = '.'
            
            for mo in moved:
                pos, val = mo
                board[pos[0]][pos[1]] = val

            board[current[0]][current[1]] = '.'
            board[current[0] + x][current[1] + y] = '@'

            current = (current[0]+x, current[1]+y)
    print_board(board)

    t = 0
    for i in range(L):
        for j in range(C):
            if board[i][j] == '[':
                t += 100 * i +j
    print(t)

# part1()
part2(board)
# print_board(board)