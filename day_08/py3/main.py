from collections import Counter


def part_one(images):
    prev = -1
    zeros = 0
    ones = 0
    twos = 0
    for _, v in images.items():
        c = Counter(v)
        curr_zeros = c.get(0, 0)
        if prev != -1:
            if curr_zeros < zeros:
                prev, zeros = zeros, curr_zeros
                ones = c.get(1, 0)
                twos = c.get(2, 0)
        else:
            prev, zeros = curr_zeros, curr_zeros
            ones = c.get(1, 0)
            twos = c.get(2, 0)
    return ones * twos


def part_two(images, width, height):
    decoded_image = ""
    size = width * height
    layers = len(images) // (width * height)

    for i in range(size):
        for j in range(layers):
            pixel = images[i + (j * size)]
            if pixel != 2:
                decoded_image += str(pixel)
                break
    return "".join(str(i) for i in decoded_image)


def print_image(image, width, height):
    start = 0
    for i in range(height):
        line = image[start : width + start]
        start += width
        print("".join(["." if j == "0" else "#" for j in line]))


if __name__ == "__main__":
    image_data = [int(i) for i in open("day_08/input.txt").read().strip()]
    width = 25
    height = 6
    layers = len(image_data) // (width * height)
    image_layers = {}
    current_pos = 0
    for i in range(layers):
        layer_data = image_data[current_pos : (width * height) + current_pos]
        current_pos += width * height
        image_layers[i] = layer_data

    print(part_one(image_layers))

    decoded_image = part_two(image_data, width, height)
    print_image(decoded_image, width, height)
