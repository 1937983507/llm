from openai import OpenAI

client = OpenAI(
    base_url = 'http://localhost:8080/v1',
    api_key='xx', # required, but unused
)

allMessages = [
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Who won the world series in 2020?"},
    {"role": "assistant", "content": "The LA Dodgers won in 2020."},
    {"role": "user", "content": "Where was it played?"}
  ]

# === 流式输出 ===
stream = client.chat.completions.create(
  model="deepseek-r1:1.5b",
  messages=allMessages,
  stream=True
)
for event in stream:
    print(event)


# # === 非流式输出 ===
# response = client.chat.completions.create(
#   model="deepseek-r1:1.5b",
#   messages=allMessages
# )
# print(response.choices[0].message.content)