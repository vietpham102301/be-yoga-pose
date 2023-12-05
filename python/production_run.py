import time
from torch import load, no_grad, softmax, max
from PIL import Image
import argparse
from torchvision import transforms
import torch.nn as nn
from torchvision.datasets import ImageFolder
from torch import device

dataset_path = r'/Users/vietpham1023/Desktop/filtered_dataset'
model_path = r'/Users/vietpham1023/Desktop/python-resource-yoga-pose/model4_30_classes_01.pth'


class ConvNet(nn.Module):
    def __init__(self, num_classes):
        super(ConvNet, self).__init__()
        self.conv1 = nn.Conv2d(in_channels=3, out_channels=32, kernel_size=3, stride=1, padding=1)
        self.relu1 = nn.ReLU()
        self.pool1 = nn.MaxPool2d(kernel_size=2, stride=2)
        self.conv2 = nn.Conv2d(in_channels=32, out_channels=64, kernel_size=3, stride=1, padding=1)
        self.relu2 = nn.ReLU()
        self.pool2 = nn.MaxPool2d(kernel_size=2, stride=2)
        self.conv3 = nn.Conv2d(in_channels=64, out_channels=128, kernel_size=3, stride=1, padding=1)
        self.relu3 = nn.ReLU()
        self.pool3 = nn.MaxPool2d(kernel_size=2, stride=2)
        self.conv4 = nn.Conv2d(in_channels=128, out_channels=256, kernel_size=3, stride=1, padding=1)
        self.relu4 = nn.ReLU()
        self.pool4 = nn.MaxPool2d(kernel_size=2, stride=2)
        self.conv5 = nn.Conv2d(in_channels=256, out_channels=512, kernel_size=3, stride=1, padding=1)
        self.relu5 = nn.ReLU()
        self.pool5 = nn.MaxPool2d(kernel_size=2, stride=2)
        self.fc1 = nn.Linear(512 * 7 * 7, 512)  # Adjust input size based on dataset
        self.relu6 = nn.ReLU()
        self.fc2 = nn.Linear(512, num_classes)

    def forward(self, x):
        x = self.conv1(x)
        x = self.relu1(x)
        x = self.pool1(x)
        x = self.conv2(x)
        x = self.relu2(x)
        x = self.pool2(x)
        x = self.conv3(x)
        x = self.relu3(x)
        x = self.pool3(x)
        x = self.conv4(x)
        x = self.relu4(x)
        x = self.pool4(x)
        x = self.conv5(x)
        x = self.relu5(x)
        x = self.pool5(x)
        x = x.view(x.size(0), -1)
        x = self.fc1(x)
        x = self.relu6(x)
        x = self.fc2(x)
        return x


def main(input_image_path):
    start_time = time.time()

    data_transforms = transforms.Compose([
        transforms.Resize((224, 224)),
        transforms.RandomHorizontalFlip(),
        transforms.ToTensor(),
        transforms.Normalize([0.485, 0.456, 0.406], [0.229, 0.224, 0.225])
    ])

    dataset = ImageFolder(root=dataset_path, transform=data_transforms)
    num_classes = len(dataset.classes)
    print(num_classes)
    mps_device = device("mps:0")
    model = ConvNet(num_classes)
    model.to(mps_device)

    model.load_state_dict(load(model_path))
    model.eval()

    image = Image.open(input_image_path)
    image = image.convert('RGB')

    image = data_transforms(image).unsqueeze(0)

    with no_grad():
        image = image.to(mps_device)
        outputs = model(image)
        probabilities = softmax(outputs, dim=1)
        _, predicted = max(outputs, 1)
        confidence = max(probabilities)

        # Get the class name corresponding to the predicted class index
        predicted_class = dataset.classes[predicted.item()]

        # Print the predicted class and confidence
        print(f'Predicted Class: {predicted_class}')
        print(f'Confidence: {confidence.item()}')

    end_time = time.time()
    execution_time = end_time - start_time
    print(f"Execution time: {execution_time} seconds")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Image classification script")
    parser.add_argument("--image", type=str, help="Path to the input image", required=True)
    args = parser.parse_args()

    main(args.image)
