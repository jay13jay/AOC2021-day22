import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
import numpy as np

# Define the coordinates of the cube
x = [10, 10, 10, 10, 10, 10, 10, 10]
y = [10, 10, 11, 11, 10, 10, 11, 11]
z = [10, 11, 11, 10, 10, 11, 11, 10]

# Create a 3D plot
fig = plt.figure()
ax = fig.add_subplot(111, projection='3d')

# Plot the cube
ax.plot(x, y, z, color='blue')

# Set labels and title
ax.set_xlabel('X')
ax.set_ylabel('Y')
ax.set_zlabel('Z')
ax.set_title('Cube at (10, 10, 10)')

# Show the plot
plt.show()
