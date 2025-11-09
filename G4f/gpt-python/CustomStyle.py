from g4f.client import Client
import g4f.Provider as Provider
import os
import urllib.parse
import shutil
import sys

client = Client(provider=Provider.OpenAIFM)
if len(sys.argv) > 1:
    user_content = sys.argv[1]
else:
    user_content = "This is amazing news! We're launching a new product that will change the industry!" # Default fallback
   
# Custom voice style
custom_instructions = """
Voice: Energetic and enthusiastic, with a youthful quality.

Tone: Upbeat and positive, conveying excitement and passion about the topic.

Pacing: Quick but clear, with dynamic variations to maintain interest.

Emotion: Genuinely excited and enthusiastic, with occasional moments of awe or surprise.

Pronunciation: Crisp and clear, with emphasis on key words to highlight important points.

Pauses: Strategic short pauses before important information, creating anticipation.
"""

response = client.media.generate(
    user_content,
    model="gpt-4o-mini-tts",
    audio={
        "voice": "nova",
        "instructions": custom_instructions
    }
)

# Decode the URL to handle encoded characters in the filename
source_path = response.data[0].url.replace("/media", "./generated_media/")
decoded_source_path = urllib.parse.unquote(source_path)
shutil.copy(decoded_source_path, "./output/tmp.mp3")
os.rename(decoded_source_path, "./output/tmp.wav")
