import matplotlib.pyplot as plt

plt.subplot(221)
yy = [2.875890502670461e-08,3.020958589988063e-08,3.6541537047034706e-08,4.4370057075236115e-08,]
y = [1/i for i in yy]
x = [0.001,0.034,0.067,0.1,]
plt.plot(x, y, "o--", label='оптимизированный')
yy = [1.534387368444713e-08,3.644823611183285e-08,5.195847968877397e-08,6.793641169000401e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный')
plt.axis((0, 1e-1+1e-3, 0, 0.7*1e+8))
plt.ylabel('ops')
# plt.xlabel('вероятность')
plt.legend()
plt.title("pw")

plt.subplot(222)
yy = [2.4044679041666667e-08,3.3299781476562494e-08,3.354815195138889e-08,3.449766951302083e-08,3.5056021728125e-08,3.5397725440104163e-08,3.597444779241072e-08,3.650882462304687e-08,3.708774463773148e-08,3.7531936166666664e-08,3.809745775804924e-08,3.859581501779514e-08,]
y = [1/i for i in yy]
x = [1,2,3,4,5,6,7,8,9,10,11,12,]
plt.plot(x, y, "o--", label='оптимизированный')
yy = [3.495895965794309e-08,4.56893024007822e-08,4.1374732245243863e-08,4.1827984513088154e-08,4.3152456894295666e-08,4.332431681569106e-08,4.361979851803294e-08,4.358977423154824e-08,4.396549587364896e-08,4.416318903265232e-08,4.453107837255879e-08,4.486391496968862e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный')
plt.axis((0, 13, 0, 0.45*1e+8))
plt.ylabel('ops')
# plt.xlabel('вероятность')
plt.legend()
plt.title("procs")

plt.subplot(223)
yy = [2.0063931764505546e-08,3.049441376344885e-08,3.9979810496195775e-08,4.9341929024705886e-08,]
y = [1/i for i in yy]
x = [1,550,1100,1650,]
plt.plot(x, y, "o--", label='оптимизированный')
yy = [2.9858769646102305e-08,3.8605128440300784e-08,4.7214131024872375e-08,5.600897206378251e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный')
plt.axis((0, 1651, 0, 0.55*1e+8))
plt.ylabel('ops')
plt.xlabel('операции')
plt.legend()
plt.title("w")

plt.subplot(224)
yy = [3.419994309868733e-08,3.366755438494537e-08,3.483628802036968e-08,3.717629954485367e-08,]
y = [1/i for i in yy]
x = [1,4,7,10,]
plt.plot(x, y, "o--", label='оптимизированный')
yy = [4.2263389887265895e-08,4.227051786181628e-08,4.2659696749457546e-08,4.449339667651825e-08,]
y = [1/i for i in yy]
plt.plot(x, y, "o--", label='распределенный')
plt.axis((0, 11, 0, 0.33*1e+8))
plt.ylabel('ops')
plt.xlabel('операции')
plt.legend()
plt.title("r")

plt.show()
