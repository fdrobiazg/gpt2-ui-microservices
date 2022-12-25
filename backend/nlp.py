from flask import Flask
from transformers import GPT2LMHeadModel , GPT2Tokenizer

class Model():
    def __init__(self):
        self.tokenizer = GPT2Tokenizer.from_pretrained('gpt2-large') 
        self.model = GPT2LMHeadModel.from_pretrained('gpt2-large' , 
                        pad_token_id = self.tokenizer.eos_token_id)
    
    def input_handler(self, input_text):
        input_ids = self.tokenizer.encode(input_text, return_tensors = 'pt')
        return self.tokenizer.decode(input_ids[0])

    def generate_text(self, processed_input):
        output = self.model.generate(processed_input, 
                                max_length = 50, 
                                num_beams = 5,
                                no_repeat_ngram_size  = 2,
                                early_stopping = True)
        return self.tokenizer.decode(output[0], skip_special_tokens = True)


app = Flask(__name__)

@app.route("/api/generate")
def generate():
    
