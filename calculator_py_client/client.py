from __future__ import print_function

import grpc
import calculator_pb2
import calculator_pb2_grpc
import sys

def run():
    if len(sys.argv) != 3:
        sys.exit("Usage: python client.py <n1> <n2>")

    n1 = int(sys.argv[1])
    n2 = int(sys.argv[2])

    with grpc.insecure_channel('localhost:50051') as channel:
        stub = calculator_pb2_grpc.CalculatorStub(channel)

        try:
            # Perform gRPC Add operation.
            add_response = stub.Add(calculator_pb2.AddRequest(n1=n1, n2=n2))
            print(f'{n1} + {n2} = {add_response.n1}')

            # Perform gRPC Subtract operation.
            subtract_response = stub.Subtract(calculator_pb2.SubtractRequest(n1=n1, n2=n2))
            print(f'{n1} - {n2} = {subtract_response.n1}')

            # Perform gRPC Multiply operation.
            multiply_response = stub.Multiply(calculator_pb2.MultiplyRequest(n1=n1, n2=n2))
            print(f'{n1} * {n2} = {multiply_response.n1}')

            # Perform gRPC Divide operation.
            divide_response = stub.Divide(calculator_pb2.DivideRequest(n1=n1, n2=n2))
            if divide_response == "Cannot divide by zero":
                print(f'Dividing error: Cannot divide by zero')
            else:
                print(f'{n1} / {n2} = {divide_response.n1:.2f}')

        except grpc.RpcError as e:
            print(f"Error: {e.code()}\n{e.details()}")

if __name__ == '__main__':
    run()