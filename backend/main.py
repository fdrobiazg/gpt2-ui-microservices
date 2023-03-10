from flask import Flask, request
from transformers import GPT2LMHeadModel , GPT2Tokenizer

class Model():
    def __init__(self):
        self.tokenizer = GPT2Tokenizer.from_pretrained('gpt2-large') 
        self.model = GPT2LMHeadModel.from_pretrained('gpt2-large' , 
                        pad_token_id = self.tokenizer.eos_token_id)
    
    def input_handler(self, input_text):
        input_ids = self.tokenizer.encode(input_text, return_tensors = 'pt')
        self.tokenizer.decode(input_ids[0])
        return input_ids

    def generate_text(self, processed_input):
        output = self.model.generate(processed_input, 
                                max_length = 50, 
                                num_beams = 5,
                                no_repeat_ngram_size  = 2,
                                early_stopping = True)
        return self.tokenizer.decode(output[0], skip_special_tokens = True)


app = Flask(__name__)

@app.route("/api/generate", methods=["GET", "POST"])
def generate():
    model = Model()
    data = str(request.data)
    print(data)
    return '.'.join(model.generate_text(model.input_handler(data)).split('.')[:-1]) + '.'


    
if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5001, debug=True)