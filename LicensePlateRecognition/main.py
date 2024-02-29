from flask import Flask, request, jsonify
from flask_cors import CORS
import easyocr

app = Flask(__name__)
CORS(app)

reader = easyocr.Reader(['ch_sim', 'en'])  # 初始化OCR reader

@app.route('/ocr', methods=['POST'])
def ocr():
    if 'file' not in request.files:
        return jsonify({'error': 'No file part'}), 400
    file = request.files['file']
    if file.filename == '':
        return jsonify({'error': 'No selected file'}), 400
    if file:
        # 将上传的图片保存到临时文件中
        temp_path = 'temp_image.png'
        file.save(temp_path)
        
        # 使用easyocr进行文本识别
        result = reader.readtext(temp_path)
        
        
        # 提取识别的文本
           # 提取识别的文本
        texts = [detection[1] for detection in result]
        return jsonify({'texts': texts}), 200
           


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8084)