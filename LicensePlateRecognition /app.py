import cv2
import numpy as np
import pytesseract

def imreadex(filename):
    return cv2.imdecode(np.fromfile(filename, dtype=np.uint8), cv2.IMREAD_COLOR)

def find_plate_area(img):
    """改进的车牌区域检测"""
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
    blur = cv2.bilateralFilter(gray, 11, 17, 17)  # 使用双边滤波而不是高斯模糊
    edged = cv2.Canny(blur, 30, 200)  # 调整Canny边缘检测参数
    contours, _ = cv2.findContours(edged.copy(), cv2.RETR_TREE, cv2.CHAIN_APPROX_SIMPLE)
    plates = []
    for c in contours:
        peri = cv2.arcLength(c, True)
        approx = cv2.approxPolyDP(c, 0.02 * peri, True)
        if len(approx) == 4:  # 寻找四边形轮廓
            x, y, w, h = cv2.boundingRect(approx)
            aspect_ratio = w / float(h)
            if 2 < aspect_ratio < 5:
                plates.append((x, y, w, h))
    return plates

def preprocess_for_ocr(img, plate_area):
    x, y, w, h = plate_area
    plate_img = img[y:y+h, x:x+w]
    gray_plate = cv2.cvtColor(plate_img, cv2.COLOR_BGR2GRAY)
    # 使用自适应阈值而不是全局阈值
    binary_plate = cv2.adaptiveThreshold(gray_plate,255,cv2.ADAPTIVE_THRESH_GAUSSIAN_C,\
                cv2.THRESH_BINARY,11,2)
    return binary_plate

def main(image_path):
    img = imreadex(image_path)
    plates = find_plate_area(img)

    recognized_texts = []
    for plate_area in plates:
        binary_plate = preprocess_for_ocr(img, plate_area)
        # 使用改进的Tesseract配置
        text = pytesseract.image_to_string(binary_plate, config='--psm 8')  # 尝试不同的PSM模式
        recognized_texts.append(text.strip())

    # 删除重复值和空字符串
    recognized_texts = list(filter(None, set(recognized_texts)))
    for text in recognized_texts:
        print("Recognized text:", text)

if __name__ == "__main__":
    image_path = "test1.png"
    main(image_path)
