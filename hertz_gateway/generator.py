import subprocess
import shlex
import random
import sys

sample_command = "curl --location --request POST http://localhost:8888/math/add --header 'Content-Type: application/json' --data '{\"first\": 1, \"second\": 0}'"
separated_command = shlex.split(sample_command)
'''
separated_command = ['curl', '--location', '--request', 'POST', 'http://localhost:8888/math/add', '--header', 'Content-Type: application/json', '--data',
                     '{"first": 1, "second": 0}']
'''

service = ["echo/echo", "math"]
math_svc = ["add", "subtract", "multiply", "divide"]

random_echo_service = ["hello", "hi", "message sending", "peppikah sending message", "yuancheng-yugan", "cloudwego"]
random_math_svc = range(-10, 11)

for i in range(int(sys.argv[2])):
    if sys.argv[1] == "echo":
        chosen_service = service[0]
        separated_command[4] = 'http://localhost:8888/' + chosen_service
        separated_command[-1] = '{"message" : "' + random.choice(random_echo_service) + '"}'
    else:
        random_math_service = random.choice(math_svc)
        chosen_math_service = service[1] + "/" + random_math_service
        separated_command[4] = 'http://localhost:8888/' + chosen_math_service
        separated_command[-1] = '{\"first\": ' + str(random.choice(random_math_svc)) + ', \"second\": ' +str(random.choice(random_math_svc)) +'}'       
    print(" ".join(separated_command))
    subprocess.run(separated_command)
    print("\n")
