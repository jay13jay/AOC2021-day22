import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D

# Define the cuboids
cuboids = [
    (-20, 26, -36, 17, -47, 7),
    (-20, 33, -21, 23, -26, 28),
    (-22, 28, -29, 23, -38, 16),
    (-46, 7, -6, 46, -50, -1),
    (-49, 1, -3, 46, -24, 28),
    (2, 47, -22, 22, -23, 27),
    (-27, 23, -28, 26, -21, 29),
    (-39, 5, -6, 47, -3, 44),
    (-30, 21, -8, 43, -13, 34),
    (-22, 26, -27, 20, -29, 19),
    (-48, -32, 26, 41, -47, -37),
    (-12, 35, 6, 50, -50, -2),
    (-48, -32, -32, -16, -15, -5),
    (-18, 26, -33, 15, -7, 46),
    (-40, -22, -38, -28, 23, 41),
    (-16, 35, -41, 10, -47, 6),
    (-32, -23, 11, 30, -14, 3),
    (-49, -5, -3, 45, -29, 18),
    (18, 30, -20, -8, -3, 13),
    (-41, 9, -7, 43, -33, 15),
]

# Create a 3D plot
fig = plt.figure()
ax = fig.add_subplot(111, projection='3d')

# Plot each cuboid as a wireframe box
for cuboid in cuboids:
    x = [cuboid[0], cuboid[1], cuboid[1], cuboid[0],
         cuboid[0], cuboid[1], cuboid[1], cuboid[0]]
    y = [cuboid[2], cuboid[2], cuboid[3], cuboid[3],
         cuboid[2], cuboid[2], cuboid[3], cuboid[3]]
    z = [cuboid[4], cuboid[4], cuboid[4], cuboid[4],
         cuboid[5], cuboid[5], cuboid[5], cuboid[5]]
    ax.plot(x, y, z, color='blue')

# Set labels and title
ax.set_xlabel('X')
ax.set_ylabel('Y')
ax.set_zlabel('Z')
ax.set_title('Cuboids')

# Show the plot
plt.show()
