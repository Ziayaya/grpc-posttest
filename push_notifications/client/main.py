import grpc
import time
from concurrent import futures
from notif_pb2 import FruitRequest, PoetryRequest
from notif_pb2_grpc import NotifServiceStub


def process_fruit_responses(responses):
    for response in responses:
        print(response)

def process_poetry_responses(responses):
    for response in responses:
        print(response)

def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = NotifServiceStub(channel)

    print("Welcome!")
    print("We have 2 service in here such as: ")
    print("1. Fruit content")
    print("You will know what nutrition and information about fruit you want")
    print("We just ")
    print("2. Poem picker based on lines")
    print("You will get english poem based on how much line they have")
    print("We just give you 1-100 row for poem request")
    choice = input(str("Please choose 1 or 2 (or 'exit' for out): "))

    if choice == 'exit':
        print("Goodbye!")
        return
    
    elif choice == '1':
        print(" ")
        fruit_input = input("Please input name fruit (Example: Pineapple,Apple,Cherry): ")
        fruit_names = [name.strip().capitalize() for name in fruit_input.split(',')]
        for name in fruit_names:
            try: 
                responses = stub.GetFruitInfo(FruitRequest(name=name))  # Hanya mengambil detail buah pertama
                process_fruit_responses(responses)
                time.sleep(2)
            except Exception as e:
                print("There's no fruit like that in our database or some error might be occured.")
                print(f"Error processing response for {name}: {e}")
                
        print("Goodbye!")

    elif choice == '2':
        line_count = int(input("Please input how much row you want to search poem (0 for out): "))
        if line_count > 0 and line_count <= 100:
            responses = [stub.GetPoetryInfo(PoetryRequest(linecount=line_count))]
            process_responses = process_poetry_responses
            for response in responses:
                try:
                    process_responses(response)
                    time.sleep(2)
                except Exception as e:
                    print(f"Error processing response: {e}")

        elif line_count > 100:
            print("Too much row")

        else:
            print("Jumlah baris puisi tidak valid.")
            return
        
        print("Goodbye!")
        

    else:
        print("Pilihan tidak valid.")
        return

    # Tutup channel gRPC
    channel.close()

if __name__ == '__main__':
    run()