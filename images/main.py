from PIL import Image
import os

with os.scandir(".") as it:
    for entry in it:
        if entry.name.endswith(".png") and entry.is_file():
            img = Image.open(entry.name)
            f = open(entry.name.replace(".png", ""), "w")

            width, height = img.size
            print('width :', width)
            print('height:', height)

            rgb_img = img.convert('RGB')

            f.write(f"{height} {width}\n")
            for y in range(height):
                for x in range(width):
                    pixel = rgb_img.getpixel((x, y))
                    # print(pixel)
                    r, g, b = pixel
                    print(x, y, f"#{r:02x}{g:02x}{b:02x}")
                    f.write(f"#{r:02x}{g:02x}{b:02x} ")
                f.write("\n")

            f.close()