import matplotlib.pyplot as plt

plt.subplot(221)
x = [0.001,0.034,0.067,0.1,]
yy = [2.9424539711549752e-08,5.232341822701475e-08,7.302704210594224e-08,9.345921458327353e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный', c='tab:orange')
plt.axis((0, 0.101, 0, 0.375*1e+8))
plt.ylabel('ops')
# plt.xlabel('pc')
plt.legend()
plt.title("pw = {0.001, 0.034, 0.067, 0.1}")

plt.subplot(222)
x = [1e-05,0.00034,0.00067,0.001,]
yy = [2.8543654882908908e-08,2.899777192516409e-08,2.9169774514693287e-08,2.9411783398557575e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный', c='tab:orange')
plt.axis((0, 0.00101, 0, 0.375*1e+8))
plt.ylabel('ops')
# plt.xlabel('pc')
plt.legend()
plt.title("pw = {0.001, 0.034, 0.067, 0.1}")

plt.subplot(223)
x = [0.001,0.034,0.067,0.1,]
yy = [1.1162171442640846e-08,2.1744532813432997e-08,3.01553421566704e-08,3.82902741302012e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный', c='tab:orange')
plt.axis((0, 0.101, 0, 0.99*1e+8))
plt.ylabel('ops')
plt.xlabel('pc')
plt.legend()
plt.title("pw = {0.00001, 0.00034, 0.00067, 0.001}")

plt.subplot(224)
x = [1e-05,0.00034,0.00067,0.001,]
yy = [1.2340620733288182e-08,1.1389349424164517e-08,1.1200279381970145e-08,1.1083601548687802e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный', c='tab:orange')
plt.axis((0, 0.00101, 0, 0.99*1e+8))
plt.ylabel('ops')
plt.xlabel('pc')
plt.legend()
plt.title("pw = {0.00001, 0.00034, 0.00067, 0.001}")

plt.show()
