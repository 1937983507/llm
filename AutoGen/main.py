from autogen import AssistantAgent, UserProxyAgent

config_list = [
  {
    "model": "deepseek-r1:8b",
    "base_url": "http://localhost:8080/v1",
    "api_key": "xxx",
  }
]

assistant = AssistantAgent("assistant", llm_config={"config_list": config_list})

user_proxy = UserProxyAgent(
    "user_proxy", 
    code_execution_config={"work_dir": "coding", "use_docker": False},
    human_input_mode="NEVER",                       # 关键参数：禁用人工输入
    max_consecutive_auto_reply=5,                   # 最大自动回复次数
    default_auto_reply="你分析的对，请继续分析"       # 自动回复 assistant
)
user_proxy.initiate_chat(assistant, message="你好请你介绍一下自己")  # Plot a chart of NVDA and TESLA stock price change YTD
