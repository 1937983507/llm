import asyncio

from openai import AsyncOpenAI
from agents import (
    Agent,
    Runner,
    set_default_openai_api,
    set_default_openai_client,
    set_tracing_disabled,
)
from openai.types.responses import ResponseTextDeltaEvent


# =========== 配置区 ===========
BASE_URL = "http://localhost:8080/v1"
API_KEY = "dummy_key"   # 本地模型一般不用验证，随便写即可
MODEL_NAME = "deepseek-r1:1.5b"
isStream = True         # 流式输出
inputMessage = "你好"   # 提问内容


# =========== 初始化全局设置 ===========
set_default_openai_api("chat_completions")
set_default_openai_client(AsyncOpenAI(base_url=BASE_URL, api_key=API_KEY))
# 在此示例中禁用追踪
# # 如需使用自定义追踪处理器，请参考：https://openai.github.io/openai-agents-python/tracing/#external-tracing-processors-list
set_tracing_disabled(disabled=True)
agent = Agent(
    name="Assistant",
    instructions="You are a helpful assistant",
    model=MODEL_NAME
)


# =========== 非流式执行 ===========
def run_sync_mode(user_input: str):
    result = Runner.run_sync(agent, user_input)
    print("\n=== Final Output ===")
    print(result.final_output)


# =========== 流式执行 ===========
async def run_stream_mode(user_input: str):
    result = Runner.run_streamed(agent, input=user_input)
    print("\n=== Streaming Output ===")
    async for event in result.stream_events():
        if event.type == "raw_response_event" and isinstance(event.data, ResponseTextDeltaEvent):
            print(event.data.delta, end="", flush=True)
    print()  # 换行


# =========== 主入口 ===========
if __name__ == "__main__":
    if isStream:
        asyncio.run(run_stream_mode(inputMessage))
    else:
        run_sync_mode(inputMessage)
