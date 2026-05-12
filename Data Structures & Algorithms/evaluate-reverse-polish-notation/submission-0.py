class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        operators = {"+", "-", "*", "/"}

        def calc(a, b, token):
            if token == "+":
                return a + b
            if token == "-":
                return a - b
            if token == "*":
                return a * b
            if token == "/":
                return int(a / b)

        for token in tokens:
            if token in operators:
                b = stack.pop()
                a = stack.pop()

                stack.append(calc(a, b, token))
            else:
                stack.append(int(token))

        return stack[-1]
