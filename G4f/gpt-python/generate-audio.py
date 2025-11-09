import asyncio
from g4f.client import AsyncClient
import g4f.Provider
import sys
async def main():
    file_name="./output/tmp.mp3"
    client = AsyncClient(provider=g4f.Provider.PollinationsAI)
    if len(sys.argv) > 1:
        user_content = sys.argv[1]
    else:
        user_content = "Say in UK accent-  How are you"  # Default fallback
    response = await client.chat.completions.create(
        model="openai-audio",
        messages=[{"role": "user", "content": user_content}],
        audio={"voice": "alloy", "format": "mp3"},
    )
    response.choices[0].message.save(file_name)
    print(file_name)

asyncio.run(main())
