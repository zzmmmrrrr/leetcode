from langchain_core.tools import tool
from langchain_openai import ChatOpenAI
from langchain_core.messages import HumanMessage, ToolMessage

# 1. 使用 @tool 装饰器直接定义工具（LangChain 会自动解析函数名、Docstring 和参数类型作为 Schema）
@tool
def get_weather(location: str) -> str:
    """获取指定城市的天气"""
    return f"{location}的气温是 25 度"

# 2. 初始化模型并通过 bind_tools 绑定工具组
model = ChatOpenAI(model="gpt-4o")
model_with_tools = model.bind_tools([get_weather])

# 3. 第一次调用：模型决定是否调用工具
messages = [HumanMessage(content="北京天气怎么样？")]
ai_msg = model_with_tools.invoke(messages)

# 4. 第二次调用：手动或通过 Agent 框架执行本地函数并喂回结果
if ai_msg.tool_calls:
    messages.append(ai_msg)  # 追加大模型的 Tool Call 提议
    
    for tool_call in ai_msg.tool_calls:
        # 模拟路由：根据模型提议的名称执行本地函数（此处简写）
        tool_output = get_weather.invoke(tool_call["args"])
        
        # 将执行结果作为 ToolMessage 追加到上下文，注意必须匹配 tool_call_id
        messages.append(ToolMessage(content=str(tool_output), tool_call_id=tool_call["id"]))
    
    # 再次调用模型，获取最终的自然语言回答
    final_response = model_with_tools.invoke(messages)
    print(final_response.content)