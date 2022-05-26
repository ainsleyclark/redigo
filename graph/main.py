import matplotlib.pyplot as plt

ns_encode = [2**x for x in range(11)];
data = {
  'JSON': [159.1, 335.7, 576.7, 737.4,  889.0, 1071, 1241, 1439, 1582,1774, 1997],
  'Gob': [958.5, 1092, 1136, 1213, 1251, 1291, 1395, 1492, 1489, 1570, 1621],
  'Message Pack': [145.5, 279.8, 386.8, 491.7, 631.8, 712.8, 807.4, 903.2, 1048, 1198, 1312],
}
for (label, values) in data.items():
    plt.plot(ns_encode, values, label=label)
plt.legend()
plt.show()

ns_decode = [2**x for x in range(11)];
data = {
  'JSON': [163.6, 514.7, 721.0, 982.1, 1215, 1439, 1680,1947, 2178, 2709, 2961],
  'Gob': [9922, 9935, 9946, 9937, 10100, 10134, 10102, 10072, 10221, 10400, 10562],
  'Message Pack': [139.6, 358.2, 525.3,  713.8, 857.3, 1035, 1238, 1421, 1516, 1848, 2010],
}
for (label, values) in data.items():
    plt.plot(ns_decode, values, label=label)
plt.legend()
plt.show()
