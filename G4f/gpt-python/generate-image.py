from g4f.client import Client
import sys
import os
client = Client()
if len(sys.argv) > 1:
    user_content = sys.argv[1]
else:
    user_content = "a white siamese cat"  # Default fallback
response = client.images.generate(
    model="flux",
    prompt=user_content,
    response_format="png"
)
source_path = response.data[0].url.replace("/media", "./generated_media/")
os.rename(source_path, "./output/tmp.png")
print(f"Generated image URL: {response.data[0]}")