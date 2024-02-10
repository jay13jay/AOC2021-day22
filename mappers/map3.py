import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
import numpy as np

# Define the ranges for the first cuboid
x_range_1 = (-20, 26)
y_range_1 = (-36, 17)
z_range_1 = (-47, 7)

# Define the ranges for the second cuboid
x_range_2 = (-20, 33)
y_range_2 = (-21, 23)
z_range_2 = (-26, 28)

# Define the vertices of the first cuboid
vertices_1 = [
    (x_range_1[0], y_range_1[0], z_range_1[0]),
    (x_range_1[1], y_range_1[0], z_range_1[0]),
    (x_range_1[1], y_range_1[1], z_range_1[0]),
    (x_range_1[0], y_range_1[1], z_range_1[0]),
    (x_range_1[0], y_range_1[0], z_range_1[1]),
    (x_range_1[1], y_range_1[0], z_range_1[1]),
    (x_range_1[1], y_range_1[1], z_range_1[1]),
    (x_range_1[0], y_range_1[1], z_range_1[1])
]

# Define the vertices of the second cuboid
vertices_2 = [
    (x_range_2[0], y_range_2[0], z_range_2[0]),
    (x_range_2[1], y_range_2[0], z_range_2[0]),
    (x_range_2[1], y_range_2[1], z_range_2[0]),
    (x_range_2[0], y_range_2[1], z_range_2[0]),
    (x_range_2[0], y_range_2[0], z_range_2[1]),
    (x_range_2[1], y_range_2[0], z_range_2[1]),
    (x_range_2[1], y_range_2[1], z_range_2[1]),
    (x_range_2[0], y_range_2[1], z_range_2[1])
]

# Define the edges of the first cuboid
edges_1 = [
    (0, 1), (1, 2), (2, 3), (3, 0),
    (4, 5), (5, 6), (6, 7), (7, 4),
    (0, 4), (1, 5), (2, 6), (3, 7)
]

# Define the edges of the second cuboid
edges_2 = [
    (0, 1), (1, 2), (2, 3), (3, 0),
    (4, 5), (5, 6), (6, 7), (7, 4),
    (0, 4), (1, 5), (2, 6), (3, 7)
]

# Create a 3D plot
fig = plt.figure()
ax = fig.add_subplot(111, projection='3d')

# Plot the first cuboid
for edge in edges_1:
    ax.plot([vertices_1[edge[0]][0], vertices_1[edge[1]][0]],
            [vertices_1[edge[0]][1], vertices_1[edge[1]][1]],
            [vertices_1[edge[0]][2], vertices_1[edge[1]][2]], color='blue')

# Plot the second cuboid
for edge in edges_2:
    ax.plot([vertices_2[edge[0]][0], vertices_2[edge[1]][0]],
            [vertices_2[edge[0]][1], vertices_2[edge[1]][1]],
            [vertices_2[edge[0]][2], vertices_2[edge[1]][2]], color='red')

# Set labels and title
ax.set_xlabel('X')
ax.set_ylabel('Y')
ax.set_zlabel('Z')
ax.set_title('Two Cuboids')

# Set limits for better visualization
ax.set_xlim(-50, 40)
ax.set_ylim(-40, 30)
ax.set_zlim(-50, 30)

# Show the plot
plt.show()
