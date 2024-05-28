import matplotlib.pyplot as plt

plt.subplot(111)
yy = [1.1239298933333334e-08,2.1653214720000002e-08,2.4179407905555557e-08,2.4958460806666664e-08,2.4938793855999998e-08,2.5260431674444446e-08,2.3879759402857144e-08,2.23797891325e-08,2.1011468726296298e-08,2.0048293024333334e-08,1.951797797969697e-08,1.9158637592500003e-08,]
y = [1/i for i in yy]
x = [1,2,3,4,5,6,7,8,9,10,11,12,]
plt.plot(x, y, "o--", label='оптимизированный')
yy = [1.1902597536666666e-08,6.457269796666666e-09,4.576243752222223e-09,3.6956058691666666e-09,3.3122796186666668e-09,3.035927848888889e-09,2.9736269876190477e-09,2.9625965829166667e-09,2.959878967777778e-09,2.9225704660000002e-09,2.925540505151515e-09,2.956339521388889e-09,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный')
plt.axis((0, 13, 0, 0.35*1e+9))
plt.ylabel('ops')
# plt.xlabel('вероятность')
plt.legend()
plt.title("procs")

plt.show()
