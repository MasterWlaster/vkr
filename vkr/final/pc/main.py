import matplotlib.pyplot as plt

plt.subplot(221)
x = [0.001,0.034,0.067,0.1,]
yy = [2.9220415984767125e-08,5.177992343018428e-08,7.259383763139079e-08,9.291548403875335e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный', c='tab:orange')
plt.axis((0, 0.101, 0, 0.4*1e+8))
plt.ylabel('ops')
# plt.xlabel('pc')
plt.legend()
plt.title("pw = {0.001, 0.034, 0.067, 0.1}")

plt.subplot(222)
x = [1e-05,0.00034,0.00067,0.001,]
yy = [2.8061185932862735e-08,2.866980583537674e-08,2.9004316715764408e-08,2.9191679143750006e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный', c='tab:orange')
plt.axis((0, 0.00101, 0, 0.4*1e+8))
plt.ylabel('ops')
# plt.xlabel('pc')
plt.legend()
plt.title("pw = {0.001, 0.034, 0.067, 0.1}")

plt.subplot(223)
x = [0.001,0.034,0.067,0.1,]
yy = [1.1254967970631687e-08,2.1155497600392695e-08,2.9789059849769203e-08,3.785400737702408e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный', c='tab:orange')
plt.axis((0, 0.101, 0, 0.99*1e+8))
plt.ylabel('ops')
plt.xlabel('pc')
plt.legend()
plt.title("pw = {0.00001, 0.00034, 0.00067, 0.001}")

plt.subplot(224)
x = [1e-05,0.00034,0.00067,0.001,]
yy = [1.2265951350255816e-08,1.1453797304468206e-08,1.1257064480716655e-08,1.1249108517588667e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный', c='tab:orange')
plt.axis((0, 0.00101, 0, 0.99*1e+8))
plt.ylabel('ops')
plt.xlabel('pc')
plt.legend()
plt.title("pw = {0.00001, 0.00034, 0.00067, 0.001}")

plt.show()
