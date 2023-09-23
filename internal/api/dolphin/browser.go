package dolphin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/Implex-ltd/engine/internal/utils"
)

var (
	token    = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiYmZkZTQ1MWJmNzEzYjQ5MzVmYTI5MjE5OGRiNWZiYjZmMTJkNjU2M2FjNmNlZmM5NTczY2RkMmMwNzEwNzBmYTJlZWQ4MDM2ODA2NDVjMTgiLCJpYXQiOjE2OTUxMTU5OTQuMTAxNzgyLCJuYmYiOjE2OTUxMTU5OTQuMTAxNzg0LCJleHAiOjE2OTc3MDc5OTQuMDkzMzA0LCJzdWIiOiIyNjc5MzU3Iiwic2NvcGVzIjpbXX0.snIlOIe2ivpXDipmFRItomqNMy_n-_4jzTtETKwya5jT5NR_Zy2Fxdl3iXOfH1by2hS-6QFu-GHW38xJkITIujaojToCtec_QS9YSevc-MoJvZHRTpDSqzmJQz3nIbDrvMSFdYWqN_auRhFFSTdi2S9Go3kz-FCa6Bhy9xC76b8fuNm22xVH-b7rWXGjJhYiw6b38dZw4yJdP2LIK_SMYoiKRwkCcYlANRCnqAw1TGPEdDB9CsUW8nrzfmp2-Npv_YQNGkoNdR47iJZZm6eaD0zRdpVqzs0nfMqOf4mRRpYiaEIkP4b9SoZofk1fULWYXYo2F8zYlm4Chw1yDUHX0RLxxMHY93S9rStGkpLC4TyRXq4A_2mVvyZ8oQmeOIppZA_WrBqAH3uo314GYWjxEHm7aiIEb96Ylt5HCQ3TcQRf1Z0xUirTbAvGF6eTQu6pOJB0gfe_wTtbCcFgWSp_R--c6C1WkL-8C6CoC2oxGW3-Elle626lrhodLWMrjkvX9otKz3iISnW3VzPqyFPt7n9lgmRn1VHDi5VN56q0pqTHvkarqydpU4ErovV_J4f2BizfFg-1XKrh5WgPjZtMmznG7g3LYAeTUkp-TqmZKcdWG2OZ4BK2W9uH0J0TBVu2QcVmVg1hL6BDsJ4ksrsr3sut58yGs_GuI_Q73wZ0zfY"
	endpoint = "http://127.0.0.1:3001/v1.0"
)

var (
	gpuData = map[string][]string{
		"Google Inc. (Intel)": {
			"ANGLE (Intel, Intel(R) HD Graphics Direct3D9Ex vs_3_0 ps_3_0, aticfx64.dll)",
			"ANGLE (Intel, Intel(R) Iris(R) Xe Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 620 Direct3D11 vs_5_0 ps_5_0, D3D11-31.0.101.2111)",
			"ANGLE (Intel, Intel(R) HD Graphics 4000 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 610 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) UHD Graphics 630 Direct3D11 vs_5_0 ps_5_0, D3D11-26.20.100.6911)",
			"ANGLE (Intel, Intel(R) HD Graphics Direct3D11 vs_5_0 ps_5_0, D3D11-10.18.10.4358)",
			"ANGLE (Intel, Intel(R) HD Graphics 630 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics P3000 Direct3D9Ex vs_3_0 ps_3_0, igdumd64.dll)",
			"ANGLE (Intel, Intel(R) UHD Graphics P630 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics Family Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) UHD Graphics 630 Direct3D11 vs_5_0 ps_5_0, D3D11-30.0.100.9684)",
			"ANGLE (Intel, Intel(R) HD Graphics 5500 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) G41 Express Chipset Direct3D9Ex vs_3_0 ps_3_0, igdumd64.dll)",
			"ANGLE (Intel, Intel(R) HD Graphics Family Direct3D9Ex vs_3_0 ps_3_0, aticfx64.dll)",
			"ANGLE (Intel, Intel(R) UHD Graphics 630 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 5300 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) UHD Graphics 730 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics Direct3D9Ex vs_3_0 ps_3_0, igdumd64.dll-9.17.10.4101)",
			"ANGLE (Intel, Intel(R) UHD Graphics 600 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 4400 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) Arc(TM) A770 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 520 Direct3D9Ex vs_3_0 ps_3_0, igdumdim64.dll)",
			"ANGLE (Intel, Intel(R) HD Graphics 400 Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (Intel, Intel(R) HD Graphics Direct3D11 vs_4_1 ps_4_1)",
			"ANGLE (Intel, Intel(R) HD Graphics Direct3D9Ex vs_3_0 ps_3_0, igdumd32.dll)",
			"ANGLE (Intel, Intel(R) UHD Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) Iris(TM) Graphics 6100 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) UHD Graphics 620 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) Q45/Q43 Express Chipset Direct3D9Ex vs_3_0 ps_3_0, igdumd64.dll)",
			"ANGLE (Intel, Intel(R) UHD Graphics 770 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 530 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Mobile Intel(R) 4 Series Express Chipset Family (Microsoft Corporation - WDDM 1.1) Direct3D9Ex vs_3_0 ps_3_0, igdumd64.dll)",
			"ANGLE (Intel, Intel(R) HD Graphics 4000 Direct3D9Ex vs_3_0 ps_3_0, aticfx64.dll)",
			"ANGLE (Intel, Intel(R) UHD Graphics 610 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 515 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 500 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 520 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 4000 Direct3D9Ex vs_3_0 ps_3_0, igdumd64.dll)",
			"ANGLE (Intel, Intel(R) HD Graphics 5000 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 4600 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 530 Direct3D11 vs_5_0 ps_5_0, D3D11-31.0.101.2111)",
			"ANGLE (Intel, Intel(R) Iris(R) Plus Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) UHD Graphics 620 Direct3D11 vs_5_0 ps_5_0, D3D11-25.20.100.6374)",
			"ANGLE (Intel, Intel(R) HD Graphics Family Direct3D9Ex vs_3_0 ps_3_0, igdumd64.dll)",
			"ANGLE (Intel, Intel(R) HD Graphics 3000 Direct3D9Ex vs_3_0 ps_3_0, igdumd32.dll)",
			"ANGLE (Intel, Intel(R) HD Graphics 3000 Direct3D9Ex vs_3_0 ps_3_0, igdumd64.dll)",
			"ANGLE (Intel, Intel(R) HD Graphics Direct3D9Ex vs_3_0 ps_3_0, igdumd64.dll)",
			"ANGLE (Intel, Intel(R) HD Graphics 3000 Direct3D11 vs_4_1 ps_4_1, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics 620 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel, Intel(R) HD Graphics Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (Intel, Intel(R) UHD Graphics 750 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (Intel(R) Iris(R) Plus Graphics Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (Intel, Intel(R) UHD Graphics 605 Direct3D11 vs_5_0 ps_5_0, D3D11)",
		},
		"Google Inc. (NVIDIA)": {
			"ANGLE (NVIDIA, NVIDIA GeForce GT 240 Direct3D11 vs_4_1 ps_4_1, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 2060 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1060 3GB Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 720 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTS 450 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 760 Direct3D9Ex vs_3_0 ps_3_0, nvd3dumx.dll)",
			"ANGLE (NVIDIA, NVIDIA Quadro K4000 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce 315M   Direct3D11 vs_4_1 ps_4_1, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1050 Ti Direct3D11 vs_5_0 ps_5_0, D3D11-31.0.15.2647)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3070 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3060 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 430 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 2070 SUPER Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 760 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3090 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 4060 Laptop GPU Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 560 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce MX350 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 980 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1070 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 770 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1650 Direct3D11 vs_5_0 ps_5_0, D3D11-31.0.15.3179)",
			"ANGLE (NVIDIA, NVIDIA Quadro M1200 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 330  Direct3D11 vs_4_0 ps_4_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1060 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce 8800 GTX Direct3D11 vs_4_0 ps_4_0)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1080 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1060 5GB Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 980 Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 650M Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (GeForce GTX 980 Direct3D11 vs_5_0 ps_5_0, D3D11-31.0.15.1694)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3090 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 440 Direct3D11 vs_5_0 ps_5_0, D3D11-23.21.13.9135)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 960 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3070 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 650 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3050 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 970 Direct3D11 vs_5_0 ps_5_0, D3D11-25.21.14.1681)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3050 OEM Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 660 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 630 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce MX130 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3050 Laptop GPU Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 420 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 630 Direct3D9Ex vs_3_0 ps_3_0, nvd3dum.dll-9.18.13.697)",
			"ANGLE (GeForce GTX 980 Direct3D11 vs_5_0 ps_5_0, D3D11-26.21.14.4274)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 4070 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA Quadro K2000D Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3060 Laptop GPU Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 440 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1650 SUPER Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce 7600 GT  Direct3D9Ex vs_3_0 ps_3_0, nvd3dumx.dll)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 970 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 1030 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA Quadro P2200 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 480 Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1060 6GB Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 950 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 980 Direct3D9Ex vs_3_0 ps_3_0)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 755M Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 635 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA NVS 510 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA Quadro K1100M Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 2070 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 980 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3070 Ti Laptop GPU Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1660 Ti Direct3D11 vs_5_0 ps_5_0, D3D11-31.0.15.3179)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 650 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3060 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1050 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 740 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1650 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 2070 with Max-Q Design Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA Quadro M2000 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1660 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 550 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 Laptop GPU Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1660 SUPER Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 2080 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 760 (192-bit) Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 950 Direct3D11 vs_5_0 ps_5_0, D3D11-10.18.13.6519)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1060 with Max-Q Design Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1650 with Max-Q Design Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA Quadro M3000M Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 750 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 2080 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce 310M        Direct3D11 vs_4_1 ps_4_1, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 670 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 4090 Laptop GPU Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1650 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1070 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA Quadro FX 2800M Direct3D11 vs_4_0 ps_4_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 640 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce 8800 GTX Direct3D11 vs_4_1 ps_4_1)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 4070 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce 9600 GT Direct3D11 vs_4_0 ps_4_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1660 Ti with Max-Q Design Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 220  Direct3D11 vs_4_1 ps_4_1, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 610 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA Quadro K2000 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce 9800 GT   Direct3D11 vs_4_0 ps_4_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 220   Direct3D11 vs_4_1 ps_4_1, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3050 Ti Laptop GPU Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 540M Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 4080 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce 210  Direct3D11 vs_4_1 ps_4_1, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 750 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1080 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 520 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 4090 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 650 Direct3D9Ex vs_3_0 ps_3_0, nvd3dumx.dll)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 2080 SUPER Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce G 105M Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 2060 SUPER Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1660 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1050 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 710 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce 310M    Direct3D11 vs_4_1 ps_4_1, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GT 730 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce RTX 3070 Laptop GPU Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (NVIDIA, NVIDIA GeForce GTX 1070 Ti Direct3D11 vs_5_0 ps_5_0, D3D11-30.0.14.7111)",
		},
		"Google Inc. (AMD)": {
			"ANGLE (AMD, Radeon RX 580 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon RX 580 2048SP Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon HD 8950 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon(TM) R4 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon (TM) Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon R7 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon HD 7800 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon RX 6600 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon RX 5500 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon (TM) R9 200 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon RX 6800 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon RX 5700 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon RX 590 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon HD 7700 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon RX 6700 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon Pro WX 3200 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon(TM) R5 240 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon R7 200 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon R5 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon(TM) R7 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon R9 200 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon HD 3200 Graphics Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (AMD, Radeon R9 200 Series Direct3D9Ex vs_3_0 ps_3_0)",
			"ANGLE (AMD, AMD Radeon HD 7900 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon (TM) R7 360 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon R5 200 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon HD 5850 Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (AMD, ATI Radeon HD 2600 Pro (Microsoft Corporation WDDM 1.1)  Direct3D9Ex vs_3_0 ps_3_0, atiumdag.dll)",
			"ANGLE (AMD, AMD Radeon RX 5600 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon(TM) Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon RX 6600 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon (TM) RX 480 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon R9 200 Series Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (AMD, Radeon Pro Vega 48 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon HD 8570 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon RX 6800 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon(TM) Vega 8 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon RX 6500 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon (TM) RX 470 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon(TM) RX Vega 11 Graphics Direct3D9Ex vs_3_0 ps_3_0, aticfx64.dll)",
			"ANGLE (AMD, AMD Radeon(TM) Vega 11 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon RX Vega Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon RX 570 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon(TM) RX 460 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon HD 8330 Direct3D11 vs_5_0 ps_5_0, D3D11-8.17.10.1433)",
			"ANGLE (AMD, Radeon RX 560 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon(TM) RX Vega 11 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon 550 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon (TM) R9 380 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon(TM) Vega 3 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, Radeon RX550/550 Series Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon RX 6650 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD, AMD Radeon(TM) RX Vega 10 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
			"ANGLE (AMD Radeon RX 5700 XT Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (AMD, AMD Radeon(TM) R2 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
		},
		"Google Inc.": {
			"Graphics",
			"ANGLE (Radeon (TM) RX 470 Graphics Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (NVIDIA GeForce GT 1030 Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (NVIDIA GeForce GTX 1660 Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (NVIDIA GeForce 210  Direct3D11 vs_4_1 ps_4_1)",
			"ANGLE (NVIDIA GeForce RTX 3060 Ti Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (NVIDIA GeForce GTX 560 SE Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (Intel(R) HD Graphics 4000 Direct3D11 vs_5_0 ps_5_0)",
			"ANGLE (NVIDIA GeForce RTX 2060 SUPER Direct3D11 vs_5_0 ps_5_0)",
			"Intel(R) HD Graphics",
			"ANGLE (Radeon RX 570 Series Direct3D11 vs_5_0 ps_5_0)",
		},
	}
)

func randomElement(slice []string) string {
	index := rand.Intn(len(slice))
	return slice[index]
}

func randomKey(m map[string][]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys[rand.Intn(len(keys))]
}

func getRandomVendorRenderer() (string, string) {
	el := randomKey(gpuData)
	gpuSlice := gpuData[el]

	if len(gpuSlice) == 0 {
		return "", ""
	}

	return el, randomElement(gpuSlice)
}

func NewDolphin(UserAgent, Os string) (*Dolphin, error) {
	Self := &Dolphin{
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
	}

	pname := "Win32"
	if Os == "mac" {
		pname = "MacIntel"
	} else {
		pname = "windows"
	}

	vendor, renderer := getRandomVendorRenderer()

	Self.Config = DolphinCreateBrowserPayload{
		Useragent: Locale{
			Mode:  "manual",
			Value: UserAgent,
		},
		//OSVersion:       "10",
		Platform:    pname,
		Name:        utils.GenerateRandomString(5),
		BrowserType: "anty", // Browser type. Available values: ['anty']. really ?????
		Webrtc: Webrtc{
			Mode: "off",
		},
		ProductSub:  "20030107",
		Vendor:      "Google Inc.",
		AppCodeName: "Mozilla",
		Webgl: Audio{
			Mode: "noise",
		},
		Canvas: Audio{
			Mode: "noise",
		},
		Timezone: Locale{
			Mode: "auto",
		},
		Locale: Locale{
			Mode: "auto",
		},
		WebglInfo: WebglInfo{
			Mode:     "manual",
			Vendor:   vendor,
			Renderer: renderer,
		},
		ClientRect: Audio{
			Mode: "off",
		},
		Memory: CPU{
			Mode:  "manual",
			Value: randomElementInt([]int64{4, 6, 8, 12, 16, 32, 64}),
		},
		CPU: CPU{
			Mode:  "manual",
			Value: randomElementInt([]int64{4, 6, 8, 12, 16, 32, 64}),
		},
		Geolocation: Geolocation{
			Mode: "auto",
		},
	}

	return Self, nil
}

func (D *Dolphin) Create() error {
	Payload, err := json.Marshal(D.Config)
	if err != nil {
		return err
	}

	fmt.Println(string(Payload))

	req, err := http.NewRequest("POST", "https://dolphin-anty-api.com/browser_profiles", bytes.NewReader(Payload))
	if err != nil {
		return err
	}

	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("content-type", "application/json")

	resp, err := D.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data DolphinCreateBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	D.UUID = fmt.Sprintf("%d", data.BrowserProfileID)
	return nil
}

func (D *Dolphin) Start() (string, error) {
	resp, err := D.Client.Get(fmt.Sprintf("%s/browser_profiles/%s/start?automation=1", endpoint, D.UUID))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data DolphinStartBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	// {"success":true,"automation":{"port":64933,"wsEndpoint":"/devtools/browser/10cd3b85-2eea-46b4-85b9-62235b17e039"}}
	return fmt.Sprintf("ws://127.0.0.1:%d%s", data.Automation.Port, data.Automation.WsEndpoint), nil
}

func (D *Dolphin) Close() error {
	D.Client.Get(fmt.Sprintf("%s/browser_profiles/%s/stop", endpoint, D.UUID))
	if err := D.Delete(); err != nil {
		return err
	}

	return nil
}

func (D *Dolphin) Delete() error {
	client := http.Client{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://dolphin-anty-api.com/browser_profiles/%s?forceDelete=1", D.UUID), nil)
	if err != nil {
		return err
	}

	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func randomElementInt(slice []int64) int64 {
	index := rand.Intn(len(slice))
	return slice[index]
}
