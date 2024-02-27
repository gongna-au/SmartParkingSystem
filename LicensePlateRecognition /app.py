import cv2
import pytesseract

# 读取图像

img = cv2.imread('test1.png')
if img is None:
    print("Error: Unable to load image. Check the file path.")
    exit()

# 灰度化
gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

# 高斯模糊
blur = cv2.GaussianBlur(gray, (7, 7), 0)

# 边缘检测
edges = cv2.Canny(blur, 100, 200)

# 使用Tesseract进行OCR识别
custom_config = r'--oem 3 --psm 6'
text = pytesseract.image_to_string(edges, config=custom_config)

print("识别的文本内容：", text)

# 注意：在实际应用中，可能还需要调整Tesseract的配置参数以获得最佳识别效果
