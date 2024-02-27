安装Flask
打开终端，运行以下命令来安装Flask。确保使用pip3而不是pip来安装，以确保库安装到Python 3而不是Python 2的环境中：
```shell
python3 --version
```

```shell
pip3 install opencv-python-headless --force-reinstall
```

```sh
pip3 install flask
```
如果你没有安装pip3，可能需要先安装它。通常，安装Python 3时pip3会自动安装。如果你是通过Homebrew（MacOS的包管理器）安装Python 3的，pip3也应该已经安装好了。

确认安装
安装完成后，你可以运行以下命令来确认Flask已正确安装：

```sh
python3 -m flask --version
```
这应该会显示Flask以及其依赖的Werkzeug库的版本信息。

重新运行你的应用
现在，再次尝试运行你的Flask应用：

```sh
python3 app.py
```
如果Flask和其他必要的库都已正确安装，你的应用应该能够在本地的8054端口启动并运行。

遇到权限问题
如果在安装过程中遇到权限问题（例如提示Permission denied），可以尝试使用sudo命令来给予安装过程管理员权限，或者使用Python的虚拟环境。虚拟环境可以让你为每个项目安装不同版本的库而不会相互冲突，使用以下命令创建和激活虚拟环境：

```sh
python3 -m venv myenv
```
```shell
source myenv/bin/activate
```
然后再次尝试安装Flask。激活虚拟环境后，使用pip安装库将默认使用该环境的pip版本，无需指定pip3。

```shell
python3.11 -m pip install --upgrade pip
```

## 安装环境

假设我们已经完成了前两步（定位文本区域和分割字符），现在需要实现第三步，字符识别。这里，使用一个预训练的模型来完成这个任务，如使用Tesseract OCR，一个开源的光学字符识别库。

首先，确保你的环境中安装了Tesseract。可以通过在终端中运行以下命令来安装Tesseract：

## 对于Ubuntu/Debian系统
```shell
sudo apt-get install tesseract-ocr
```

## 对于Mac
```shell
brew install tesseract
```

## 对于Windows，下载安装包

```shell
https://github.com/UB-Mannheim/tesseract/wiki
```

## 需要的库

```shell
pip3 install pytesseract
```


## 车牌识别的逻辑

定位文本区域：在边缘检测后的图像中定位可能的文本区域。涉及到查找特定的形状或图案，例如车牌的矩形区域。

分割字符：将这些文本区域进一步分割成单独的字符。需要使用一些形态学操作，如膨胀和腐蚀，以帮助分隔粘在一起的字符。

字符识别：对每个字符进行识别。这通常通过使用训练有素的机器学习模型来完成，例如支持向量机(SVM)、卷积神经网络(CNN)或其他图像识别算法。




## 代码修改

```shell
which tesseract
```