from pathlib import Path


def create_folders():
    """Creates folders for each AOC day up to 25th."""
    folder_names = [f"day_{d}" for d in range(4, 26, 1)]
    for folder in folder_names:
        folder = Path(folder)
        folder.mkdir(exist_ok=True)
        txt_file = folder / Path("input.txt")
        txt_file.touch(exist_ok=True)

        go = folder / Path("go")
        go_file = go / Path("main.go")
        go.mkdir(exist_ok=True, parents=True)
        go_file.touch(exist_ok=True)

        py3 = folder / Path("py3")
        py3_file = py3 / Path("main.py")
        py3.mkdir(exist_ok=True, parents=True)
        py3_file.touch(exist_ok=True)


if __name__ == "__main__":
    create_folders()
