import numpy as np
import matplotlib.pyplot as plt

#Activation function
# Helper function to plot activation function
def plot_func(x,y,title):
    pl.plot(x,y)
    plt.title(title)
    plt.xlabel('x')
    plt.ylabel('activation(x)')
    plt.grid(True)
    plt.show()

def sigmoid(x):
    return 1 / ((1 + np.exp(-1)))

def  relu(x):
    return np.maximum(0,x)

def leaky_relu(x,alpha=0.1):
    return np.maximum(alpha*x,x)

def tanh(x):
    return np.tanh(x)

def softmax(x):
    exp_scores = np.exp(x)
    return exp_scores / np.sum(exp_scores)


x = np.linspace(-10,10,100)

# a lmbda function that calculate s prediction based on the input x using
# predifined values for the weights w1 and bias b
# @TODO So where do we get the weights and given the dataset how do you calculate the weight?
prediction = lambda x,w1=.2, b=1.99: x * w1 + b
layer1_1 = np.maximum(0,prediction(x))
plt.plot(x,layer1_1)


layer1_2 = np.maximum(0,prediction(x,.3,-2))
plt.plot(x,layer1_1+layer1_2)

layer1_3 = np.maximum(0,prediction(x,.6,-2))
plt.plot(x,layer1_1+layer1_2+layer1_3)
