from dataclasses import dataclass, field
from typing import List, Optional

DIR_TYPE = "dir"
FILE_TYPE = "file"
TYPES = {DIR_TYPE, FILE_TYPE}
CMD_LS = "ls"
CMD_CD = "cd"

ROOT_PATH = "/"
BACK_PATH = ".."

TOTAL_DISK_SPACE = 70000000
UNUSED_SPACE_REQUIRED = 30000000


@dataclass
class Node:
    name: str
    type: str
    parent: Optional['Node']
    # Safe defaults
    size: int = 0
    children: List['Node'] = field(default_factory=list)

    def set_size(self, size):
        self.size = size

    def add_children(self, node: 'Node'):
        self.children.append(node)


class FileSystem:
    def __init__(self):
        self.head: Node = Node(name=ROOT_PATH, type=DIR_TYPE, parent=None)
        self.curr_dir: Optional[Node] = None

    def process_cd(self, cd_path: str):
        if self.curr_dir is None:
            if cd_path != ROOT_PATH:
                raise Exception(f"Curr dir is None, but path given is {cd_path}")
            self.curr_dir = self.head
            return

        if path == ROOT_PATH:
            self.curr_dir = self.head
            return

        if path == BACK_PATH:
            if self.curr_dir.parent is None:
                raise Exception("Can't go back from root")
            self.curr_dir = self.curr_dir.parent
            return

        for node in self.curr_dir.children:
            if node.name == path:
                self.curr_dir = node
                return

        raise Exception(f'cd {path} failed, since {self.curr_dir.name} does not have the path within it')

    def process_ls(self, contents_list: List[str]):
        for content in contents_list:
            if content.startswith("dir"):
                dir_node = Node(
                    name=content.strip().split(" ")[1],
                    type=DIR_TYPE,
                    parent=self.curr_dir
                )
                self.curr_dir.add_children(dir_node)
            else:
                size, name = content.strip().split(" ")
                child_node = Node(
                    name=name,
                    type=FILE_TYPE,
                    parent=self.curr_dir,
                    size=int(size)
                )
                self.curr_dir.add_children(child_node)

        # Recalculate size for the curr-dir and propogate upwards
        self._recalculate_dir(self.curr_dir)

    @classmethod
    def _recalculate_dir(cls, node):
        if node is None:
            return
        node.size = sum([child.size for child in node.children])
        cls._recalculate_dir(node.parent)

    def print_structure(self):
        self.__print(self.head)

    def __print(self, node: Node, level: int = 0):
        if node is None:
            return
        tabs = " " * level * 4
        print(f"{tabs}- {node.name} ({node.type}, size={node.size})")
        for child in node.children:
            self.__print(child, level + 1)

    def get_dir_sizes(self) -> List[int]:
        return self.__get_dir_sizes(self.head)

    def __get_dir_sizes(self, node) -> List[int]:
        if node is None or node.type == FILE_TYPE:
            return []
        sizes = [node.size]
        for child in node.children:
            if child.type == DIR_TYPE:
                sizes.extend(self.__get_dir_sizes(child))
        return sizes


if __name__ == "__main__":
    fs = FileSystem()
    with open("ip2.txt", "r") as fp:
        curr_command = None
        ls_buffer = []
        for curr_line in fp:
            line = curr_line.strip()
            if line.startswith("$ "):
                # Handle previous ones
                if curr_command == CMD_LS:
                    fs.process_ls(contents_list=ls_buffer)
                ls_buffer = []
                curr_command = None

                if line == "$ ls":
                    curr_command = CMD_LS
                else:
                    curr_command = CMD_CD
                    path = line.split(" ")[2]
                    fs.process_cd(path)
            else:
                ls_buffer.append(line)
        if curr_command == CMD_LS:
            fs.process_ls(contents_list=ls_buffer)

    # fs.print_structure()
    dir_sizes = fs.get_dir_sizes()
    print(dir_sizes)
    # print(sum([size for size in dir_sizes if size <= 100000]))

    curr_unused_space = UNUSED_SPACE_REQUIRED - (TOTAL_DISK_SPACE - dir_sizes[0])
    print(f"curr-unused-space {curr_unused_space}")
    print(sorted(dir_sizes))
    for size in sorted(dir_sizes):
        if size > curr_unused_space:
            print(size)
            break
