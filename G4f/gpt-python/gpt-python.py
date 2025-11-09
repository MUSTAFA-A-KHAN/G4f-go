import sys
from g4f.client import AsyncClient
import asyncio

async def main():
    # Get the text input from command-line arguments
    if len(sys.argv) > 1:
        user_content = sys.argv[1]
    else:
        user_content = "Explain quantum computing briefly"  # Default fallback

    client = AsyncClient()
    response = await client.chat.completions.create(
        model="gpt-4o-mini",
        messages=[{"role": "user", "content": user_content}],
    )
    print(response.choices[0].message.content)

asyncio.run(main())
