package gologin

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
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

func GetFingerprint() ([]byte, error) {
	vendor, renderer := getRandomVendorRenderer()

	data, err := json.Marshal(&Fingerprint{
		NewTabPage: NewTabPage{
			PrevNavigationTime: "13338424201908559",
		},
		AccountTrackerServiceLastUpdate: "13338423195083573",
		ACKExistingNTPExtensions:        true,
		AlternateErrorPages: AlternateErrorPages{
			Backup: true,
		},
		AnnouncementNotificationServiceFirstRunTime: "13337882905385312",
		Apps: FingerprintApps{
			ShortcutsArch:    "",
			ShortcutsVersion: 0,
		},
		Autocomplete: Autocomplete{
			RetentionPolicyLastVersion: 113,
		},
		Autofill: Autofill{
			OrphanRowsRemoved: true,
		},
		BookmarkBar: BookmarkBar{
			ShowOnAllTabs: false,
		},
		BookmarkEditor: BookmarkEditor{
			ExpandedNodes: []interface{}{},
		},
		Bookmarks: Bookmarks{
			EditingEnabled: true,
		},
		Browser: Browser{
			EnableSpellchecking: false,
			HasSeenWelcomePage:  false,
			WindowPlacement: WindowPlacement{
				Bottom:         1258,
				Left:           740,
				Maximized:      false,
				Right:          2036,
				Top:            404,
				WorkAreaBottom: 1392,
				WorkAreaLeft:   0,
				WorkAreaRight:  3440,
				WorkAreaTop:    0,
			},
		},
		CachedFonts: CachedFonts{
			SearchResultsPage: SearchResultsPage{
				Fallback: []interface{}{},
				Primary: []string{
					"Arial",
				},
			},
		},
		CountryidAtInstall:       21077,
		CredentialsEnableService: true,
		CustomLinks: CustomLinks{
			Initialized: true,
			List: []List{
				{
					IsMostVisited: false,
					Title:         "Facebook",
					URL:           "https://www.facebook.com/",
				},
				{
					IsMostVisited: false,
					Title:         "Google Ads",
					URL:           "https://ads.google.com/",
				},
				{
					IsMostVisited: false,
					Title:         "TikTok",
					URL:           "https://www.tiktok.com/",
				},
				{
					IsMostVisited: false,
					Title:         "Amazon",
					URL:           "https://www.amazon.com/",
				},
				{
					IsMostVisited: false,
					Title:         "eBay",
					URL:           "https://ebay.com/",
				},
				{
					IsMostVisited: false,
					Title:         "YouTube",
					URL:           "https://www.youtube.com/",
				},
				{
					IsMostVisited: false,
					Title:         "Coinlist",
					URL:           "https://coinlist.co/",
				},
				{
					IsMostVisited: false,
					Title:         "Huobi",
					URL:           "https://www.huobi.com/",
				},
				{
					IsMostVisited: false,
					Title:         "bet365",
					URL:           "https://www.bet365.com/",
				},
				{
					IsMostVisited: false,
					Title:         "PayPal",
					URL:           "https://paypal.com/",
				},
			},
		},
		DefaultAppsInstallState: 3,
		Devtools: Devtools{
			ADBKey: "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCf/Q5735xmxFKgZtUSyr67NyGzvIiZfT6ls+lj9kvwR5WvWsSP7FxTBVebe+V0Bjfv8Wbqn++lsIXYDnh4JvEWIiY5r/aQrOUVAvzNA2nUk266BhpRGTLyZJ7TArTa9VY3XZlKnyXtPhGzI1Dj1lTQey4kOcdHaw70mrLKAIzWku2Pz84yBF0fWm0ne/hcweCJVdFSEe06+8hATNm1Z9c4sbfNVbA/0blSOInPrqD/ZsWus40kOIbcg7LuF/ylJVxkmeCRQ0Vuh32GtyLrZBbG0cHp2qE4aXx4MgyOdKkEaeuo4TCHKXuYSGkNLP16yq3QlELSfoGVajTqeDKt7QaLAgMBAAECggEAGsOXLjLYs7ppBitdvsCLCr+s1CgYg/DXT5oa052q033uNQEOu/Ly+PJ+VuTxnCjEJ9AvAtYRnZq7sUqubY86KpU45tKggVBjHFI9H0DQbf7h+Z22vlVkV/dsOGmN/sknDKkMUVbxXdBR3rFsQUkkNCI/NFE0HKb0wAe3HCLebLCwrhi5IKvfdGS4q+X4myV4nCJlYJ1JJyz14B7xbfbYxHHoNwHQ91O6NhMT60KC308lJRXugcWUufxO3fbm3y3WbxMgD5vEIKTjFmt8cYmnO6hXoVTR+zcRGqnlo8Bfr9V6kCtjrz2XqBmLI0Jv7osciTmvsz1QiBVOQnwLTRe9EQKBgQDdqtLoLvgFpoZeuU5DBn+Xv6MdMTOoRedGMmCRLsSVyoCmf9j2x8f8tPo98cmOdmIF5ukG6SIKKxqgtE80mEuS7VhopYb4KvqXbUZBiWfB1PneHy5BiKkZ/S3b5hl0gta/434/5mqiPMh/CiGYgxCTJFSo0Bm6yoh+aTKBfySlMQKBgQC4xKZtbBbPsiwewtTWXPJrrquI6yixIMn/Ds2Hp8bh8dMLU7iSWpy8s1ln9ftwVlMpA7G8Zgbdpzeeb/ZRc93w0CAQE7cT2Qm/Oc13mBX5xzo49SALHHsJ7mDHOl1jyRS+Yi8qDgGfEiDufPFIzssFAfY8qysqTULpbMx2fSIoewKBgD7qIixmUF0CM7ck7zKU4JgIe1powZmsJYVF/SSX2pINRNB9Du6UK43tK+1y99sofI/iRxeACfpK26brJmPuKd0M4n/pI1LCmBdsm3k0b6+TLExAG1fcWuqWcnKsb2MjU6JpBkdKtc9QYzOzg2+PG6A3upOjZPogS/ddEkmw9YyxAoGBAIETZ8xzQTBRIP9QTk1+WyAcwtqApUmwRUy4M56793UqcautDfi38lzk0eWNctzUgX3ahdS+VbeMtI8rWiEm13n+exegHjpwWZ3MLJyKk3eN1T+rBnqcp9uVOuiS/mDSPXqOrrIrXLiaeoUApbKePDfAWhuh3CLOHJtjFScFXdJ3AoGBALFJ4sXFAZnFkgVXmoeApgX2j9jmBu71BYxTqzv5FAKD8KwJYW63uan//Qa9+lsgqFOESFAVPpd7I58KrZQJBzKAQCeUU6cA8IMZJXt++NcNYYrV5izL1crGUEAVByhkehPrCdl73s8hLuObBzAKoFKVQymZPkQH4HPn2pk8eTIX",
			Preferences: Preferences{
				InspectorDrawerSplitViewState:              "{\"horizontal\":{\"size\":0,\"showMode\":\"Both\"}}",
				InspectorViewSplitViewState:                "{\"vertical\":{\"size\":0}}",
				StylesPaneSidebarTabOrder:                  "{\"Styles\":10,\"Computed\":20}",
				CloseableTabs:                              "{\"security\":true,\"network.blocked-urls\":true}",
				ConsoleSidebarWidth:                        "{\"vertical\":{\"size\":0,\"showMode\":\"OnlyMain\"}}",
				ConsoleSidebarSelectedFilter:               "\"message\"",
				CurrentDockState:                           "\"right\"",
				DataGridCookiesTableColumnWeights:          "{\"name\":24,\"value\":34,\"domain\":7,\"path\":7,\"expires\":7,\"size\":7,\"httpOnly\":7,\"secure\":7,\"sameSite\":7,\"partitionKey\":7,\"priority\":7}",
				DrawerViewTabOrder:                         "{\"console-view\":10,\"network.blocked-urls\":20}",
				ElementsStylesSidebarWidth:                 "{\"vertical\":{\"size\":0,\"showMode\":\"OnlyMain\"}}",
				InspectorVersion:                           "35",
				NetworkBlockedPatterns:                     "[{\"enabled\":true,\"url\":\"discord.com/api/v9/users/@me?with_analytics_token=true\"}]",
				NetworkPanelSidebarState:                   "{\"vertical\":{\"size\":0,\"showMode\":\"OnlyMain\"}}",
				NetworkPanelSplitViewState:                 "{\"vertical\":{\"size\":0}}",
				NetworkPanelSplitViewWaterfall:             "{\"vertical\":{\"size\":0}}",
				NetworkResourceTypeFilters:                 "{\"XHR and Fetch\":true}",
				NetworkTextFilter:                          "\"\"",
				PanelSelectedTab:                           "\"resources\"",
				RequestInfoFormDataCategoryExpanded:        "true",
				RequestInfoGeneralCategoryExpanded:         "true",
				RequestInfoQueryStringCategoryExpanded:     "true",
				RequestInfoRequestHeadersCategoryExpanded:  "true",
				RequestInfoRequestPayloadCategoryExpanded:  "true",
				RequestInfoResponseHeadersCategoryExpanded: "true",
				ResourceViewTab:                            "\"preview\"",
				ResourcesCookiesExpanded:                   "true",
				ResourcesLastSelectedElementPath:           "[\"cookies://https://discord.com\",\"category://Cookies\"]",
				ResourcesLocalStorageExpanded:              "true",
				SourcesPanelNavigatorSplitViewState:        "{\"vertical\":{\"size\":0,\"showMode\":\"Both\"}}",
				SourcesPanelSplitViewState:                 "{\"vertical\":{\"size\":0,\"showMode\":\"Both\"},\"horizontal\":{\"size\":0,\"showMode\":\"Both\"}}",
				UndefinedTabOrder:                          "{\"sources.scopeChain\":10,\"sources.watch\":20}",
			},
			SyncedPreferencesSyncDisabled: SyncedPreferencesSyncDisabled{
				AdornerSettings: "[{\"adorner\":\"grid\",\"isEnabled\":true},{\"adorner\":\"flex\",\"isEnabled\":true},{\"adorner\":\"ad\",\"isEnabled\":true},{\"adorner\":\"scroll-snap\",\"isEnabled\":true},{\"adorner\":\"container\",\"isEnabled\":true},{\"adorner\":\"slot\",\"isEnabled\":true},{\"adorner\":\"top-layer\",\"isEnabled\":true},{\"adorner\":\"reveal\",\"isEnabled\":true}]",
				ColorFormat:     "\"original\"",
			},
		},
		DipsTimerLastUpdate: "13338423195033904",
		DomainDiversity: DomainDiversity{
			LastReportingTimestamp: "13338423195084226",
		},
		Extensions: Extensions{
			Alerts: Alerts{
				Initialized: true,
			},
			ChromeURLOverrides: TranslateSiteBlacklistWithTime{},
			Commands: map[string]Command{
				"windows:Ctrl+Shift+F": {
					CommandName: "humanTyping",
					Extension:   "lmccpkjihdfknjbpkilmbinhljdkcbak",
					Global:      false,
				},
				"windows:Ctrl+Shift+Q": {
					CommandName: "automationTask",

					Extension: "lmccpkjihdfknjbpkilmbinhljdkcbak",
					Global:    false,
				},
			},
			LastChromeVersion: "113.0.5672.92",
			Settings:          ExtensionsSettings{},
		},
		GaiaCookie: GaiaCookie{
			ChangedTime:          1658333661.108506,
			Hash:                 "2jmj7l5rSw0yVb/vlWAYkK/YBwk=",
			LastListAccountsData: "[\"gaia.l.a.r\",[]]",
		},
		Gcm: Gcm{
			ProductCategoryForSubtypes: "com.orbita.macosx",
		},
		Gologin: Gologin{
			AudioContext: AudioContext{
				Enable:     true,
				NoiseValue: rand.Float64(),
			},
			CanvasMode:             "noise",
			CanvasNoise:            rand.Float64(),
			ClientRectsNoiseEnable: false,
			DeviceMemory:           8192,
			DNS:                    "",
			DoNotTrack:             false,
			GeoLocation: GeoLocation{
				Accuracy:  100,
				Latitude:  47.8573,
				Longitude: 3.9769,
				Mode:      "prompt",
			},
			GetClientRectsNoice: rand.Float64(),
			GetClientRectsNoise: rand.Float64(),
			HardwareConcurrency: randomElementInt([]int64{4, 6, 8, 12, 16, 32, 64}),
			Icon: Icon{
				Avatar: Avatar{
					Enabled: true,
				},
				Text: generateRandomMD5(),
			},
			IsM1:       false,
			LangHeader: "fr-FR,fr;q=0.9",
			Languages:  "fr-FR,fr",
			MediaDevices: MediaDevices{
				AudioInputs:  2,
				AudioOutputs: 1,
				Enable:       true,
				Uid:          randomHex(58),
				VideoInputs:  1,
			},
			Mobile: Mobile{
				DeviceScaleFactor: 1.00000001,
				Enable:            false,
				Height:            1080,
				Width:             1920,
			},
			Name: generateRandomMD5(),
			Navigator: Navigator{
				MaxTouchPoints: 0,
				Platform:       "Win32",
			},
			Plugins: Plugins{
				AllEnable:   true,
				FlashEnable: false,
			},
			ProfileID: "64ef5eb5b342d95f657e0815",
			Proxy: Proxy{
				Password: "",
				Username: "",
			},

			ScreenHeight: 1080,
			ScreenWidth:  1920,
			StartupURL:   "",
			StartupUrls: []string{
				"",
			},
			Storage: Storage{
				Enable: true,
			},
			Timezone: Timezone{
				ID: "Europe/Paris",
			},
			UnpinableExtensionNames: []string{
				"passwords-ext",
			},
			UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
			WebGl: WebGl{
				Mode:     true,
				Renderer: renderer,
				Vendor:   vendor,
			},
			WebRTC: WebRTC{
				FillBasedOnIP:  false,
				LocalIPMasking: true,
				Mode:           "public",
				PublicIP:       generateRandomIP(),
			},
			Webgl: Webgl{
				Metadata: WebGl{
					Mode:     true,
					Renderer: renderer,
					Vendor:   vendor,
				},
			},
			WebglNoiceEnable: true,
			WebglNoiseValue:  rand.Float64(),
			WebglParams: WebglParams{
				Antialiasing: true,
				Extensions: []string{
					"EXT_color_buffer_float",
					"EXT_color_buffer_half_float",
					"EXT_disjoint_timer_query_webgl2",
					"EXT_float_blend",
					"EXT_texture_compression_bptc",
					"EXT_texture_compression_rgtc",
					"EXT_texture_filter_anisotropic",
					"EXT_texture_norm16",
					"KHR_parallel_shader_compile",
					"OES_draw_buffers_indexed",
					"OES_texture_float_linear",
					"OVR_multiview2",
					"WEBGL_compressed_texture_s3tc",
					"WEBGL_compressed_texture_s3tc_srgb",
					"WEBGL_debug_renderer_info",
					"WEBGL_debug_shaders",
					"WEBGL_lose_context",
					"WEBGL_multi_draw",
					"WEBGL_provoking_vertex",
				},
				GlCanvas: "webgl2",
				GlParamValues: []GlParamValue{
					{
						Name:  "ALIASED_LINE_WIDTH_RANGE",
						Value: map[string]int64{"0": 1, "1": 1},
					},
					{
						Name:  "ALIASED_POINT_SIZE_RANGE",
						Value: map[string]int64{"0": 1, "1": 64},
					},
					{
						Name:  []string{"DEPTH_BITS", "STENCIL_BITS"},
						Value: "n/a",
					},
					{
						Name:  "MAX_3D_TEXTURE_SIZE",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_ARRAY_TEXTURE_LAYERS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_COLOR_ATTACHMENTS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_COMBINED_TEXTURE_IMAGE_UNITS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_COMBINED_UNIFORM_BLOCKS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_CUBE_MAP_TEXTURE_SIZE",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_DRAW_BUFFERS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_FRAGMENT_INPUT_COMPONENTS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_FRAGMENT_UNIFORM_BLOCKS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_FRAGMENT_UNIFORM_COMPONENTS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_FRAGMENT_UNIFORM_VECTORS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_PROGRAM_TEXEL_OFFSET",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_RENDERBUFFER_SIZE",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_SAMPLES",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_TEXTURE_IMAGE_UNITS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_TEXTURE_LOD_BIAS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_TEXTURE_SIZE",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_UNIFORM_BLOCK_SIZE",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_UNIFORM_BUFFER_BINDINGS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_VARYING_COMPONENTS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_VARYING_VECTORS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_VERTEX_ATTRIBS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_VERTEX_OUTPUT_COMPONENTS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_VERTEX_TEXTURE_IMAGE_UNITS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_VERTEX_UNIFORM_BLOCKS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_VERTEX_UNIFORM_COMPONENTS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_VERTEX_UNIFORM_VECTORS",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "MAX_VIEWPORT_DIMS",
						Value: map[string]int64{"0": 21826, "1": 82568},
					},
					{
						Name:  "MIN_PROGRAM_TEXEL_OFFSET",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  []string{"RED_BITS", "GREEN_BITS", "BLUE_BITS", "ALPHA_BITS"},
						Value: "n/a",
					},
					{
						Name:  "RENDERER",
						Value: "WebKit WebGL",
					},
					{
						Name:  "SHADING_LANGUAGE_VERSION",
						Value: "WebGL GLSL ES 3.00 (OpenGL ES GLSL ES 3.0 Chromium)",
					},
					{
						Name:  "UNIFORM_BUFFER_OFFSET_ALIGNMENT",
						Value: rand.Intn(100000) + 1,
					},
					{
						Name:  "VENDOR",
						Value: "WebKit",
					},
					{
						Name:  "VERSION",
						Value: "WebGL 2.0 (OpenGL ES 3.0 Chromium)",
					},
				},
				ShaiderPrecisionFormat: "highp/highp",
				SupportedFunctions: []SupportedFunction{
					{
						Name:      "beginQuery",
						Supported: true,
					},
					{
						Name:      "beginTransformFeedback",
						Supported: true,
					},
					{
						Name:      "bindBufferBase",
						Supported: true,
					},
					{
						Name:      "bindBufferRange",
						Supported: true,
					},
					{
						Name:      "bindSampler",
						Supported: true,
					},
					{
						Name:      "bindTransformFeedback",
						Supported: true,
					},
					{
						Name:      "bindVertexArray",
						Supported: true,
					},
					{
						Name:      "blitFramebuffer",
						Supported: true,
					},
					{
						Name:      "clearBufferfi",
						Supported: true,
					},
					{
						Name:      "clearBufferfv",
						Supported: true,
					},
					{
						Name:      "clearBufferiv",
						Supported: true,
					},
					{
						Name:      "clearBufferuiv",
						Supported: true,
					},
					{
						Name:      "clientWaitSync",
						Supported: true,
					},
					{
						Name:      "compressedTexImage3D",
						Supported: true,
					},
					{
						Name:      "compressedTexSubImage3D",
						Supported: true,
					},
					{
						Name:      "copyBufferSubData",
						Supported: true,
					},
					{
						Name:      "copyTexSubImage3D",
						Supported: true,
					},
					{
						Name:      "createQuery",
						Supported: true,
					},
					{
						Name:      "createSampler",
						Supported: true,
					},
					{
						Name:      "createTransformFeedback",
						Supported: true,
					},
					{
						Name:      "createVertexArray",
						Supported: true,
					},
					{
						Name:      "deleteQuery",
						Supported: true,
					},
					{
						Name:      "deleteSampler",
						Supported: true,
					},
					{
						Name:      "deleteSync",
						Supported: true,
					},
					{
						Name:      "deleteTransformFeedback",
						Supported: true,
					},
					{
						Name:      "deleteVertexArray",
						Supported: true,
					},
					{
						Name:      "drawArraysInstanced",
						Supported: true,
					},
					{
						Name:      "drawBuffers",
						Supported: true,
					},
					{
						Name:      "drawElementsInstanced",
						Supported: true,
					},
					{
						Name:      "drawRangeElements",
						Supported: true,
					},
					{
						Name:      "endQuery",
						Supported: true,
					},
					{
						Name:      "endTransformFeedback",
						Supported: true,
					},
					{
						Name:      "fenceSync",
						Supported: true,
					},
					{
						Name:      "framebufferTextureLayer",
						Supported: true,
					},
					{
						Name:      "getActiveUniformBlockName",
						Supported: true,
					},
					{
						Name:      "getActiveUniformBlockParameter",
						Supported: true,
					},
					{
						Name:      "getActiveUniforms",
						Supported: true,
					},
					{
						Name:      "getBufferSubData",
						Supported: true,
					},
					{
						Name:      "getFragDataLocation",
						Supported: true,
					},
					{
						Name:      "getIndexedParameter",
						Supported: true,
					},
					{
						Name:      "getInternalformatParameter",
						Supported: true,
					},
					{
						Name:      "getQuery",
						Supported: true,
					},
					{
						Name:      "getQueryParameter",
						Supported: true,
					},
					{
						Name:      "getSamplerParameter",
						Supported: true,
					},
					{
						Name:      "getSyncParameter",
						Supported: true,
					},
					{
						Name:      "getTransformFeedbackVarying",
						Supported: true,
					},
					{
						Name:      "getUniformBlockIndex",
						Supported: true,
					},
					{
						Name:      "getUniformIndices",
						Supported: true,
					},
					{
						Name:      "invalidateFramebuffer",
						Supported: true,
					},
					{
						Name:      "invalidateSubFramebuffer",
						Supported: true,
					},
					{
						Name:      "isQuery",
						Supported: true,
					},
					{
						Name:      "isSampler",
						Supported: true,
					},
					{
						Name:      "isSync",
						Supported: true,
					},
					{
						Name:      "isTransformFeedback",
						Supported: true,
					},
					{
						Name:      "isVertexArray",
						Supported: true,
					},
					{
						Name:      "pauseTransformFeedback",
						Supported: true,
					},
					{
						Name:      "readBuffer",
						Supported: true,
					},
					{
						Name:      "renderbufferStorageMultisample",
						Supported: true,
					},
					{
						Name:      "resumeTransformFeedback",
						Supported: true,
					},
					{
						Name:      "samplerParameterf",
						Supported: true,
					},
					{
						Name:      "samplerParameteri",
						Supported: true,
					},
					{
						Name:      "texImage3D",
						Supported: true,
					},
					{
						Name:      "texStorage2D",
						Supported: true,
					},
					{
						Name:      "texStorage3D",
						Supported: true,
					},
					{
						Name:      "texSubImage3D",
						Supported: true,
					},
					{
						Name:      "transformFeedbackVaryings",
						Supported: true,
					},
					{
						Name:      "uniform1ui",
						Supported: true,
					},
					{
						Name:      "uniform1uiv",
						Supported: true,
					},
					{
						Name:      "uniform2ui",
						Supported: true,
					},
					{
						Name:      "uniform2uiv",
						Supported: true,
					},
					{
						Name:      "uniform3ui",
						Supported: true,
					},
					{
						Name:      "uniform3uiv",
						Supported: true,
					},
					{
						Name:      "uniform4ui",
						Supported: true,
					},
					{
						Name:      "uniform4uiv",
						Supported: true,
					},
					{
						Name:      "uniformBlockBinding",
						Supported: true,
					},
					{
						Name:      "uniformMatrix2x3fv",
						Supported: true,
					},
					{
						Name:      "uniformMatrix2x4fv",
						Supported: true,
					},
					{
						Name:      "uniformMatrix3x2fv",
						Supported: true,
					},
					{
						Name:      "uniformMatrix3x4fv",
						Supported: true,
					},
					{
						Name:      "uniformMatrix4x2fv",
						Supported: true,
					},
					{
						Name:      "uniformMatrix4x3fv",
						Supported: true,
					},
					{
						Name:      "vertexAttribDivisor",
						Supported: true,
					},
					{
						Name:      "vertexAttribI4i",
						Supported: true,
					},
					{
						Name:      "vertexAttribI4iv",
						Supported: true,
					},
					{
						Name:      "vertexAttribI4ui",
						Supported: true,
					},
					{
						Name:      "vertexAttribI4uiv",
						Supported: true,
					},
					{
						Name:      "vertexAttribIPointer",
						Supported: true,
					},
					{
						Name:      "waitSync",
						Supported: true,
					},
				},
				TextureMaxAnisotropyEXT: 16,
			},
			GologinWebglNoiceEnable: true,
			WebglNoiseEnable:        true,
			GologinWebglNoiseValue:  rand.Float64(),
		},
		Google: Google{
			Services: Services{
				ConsentedToSync:      false,
				SigninScopedDeviceID: uuid.NewString(),
			},
		},
		History: History{
			SavingDisabled: false,
		},
		Intl: Intl{
			AcceptLanguages:   "fr-FR,fr",
			SelectedLanguages: "fr-FR,fr",
		},
		Invalidation: Invalidation{
			PerSenderTopicsToHandler: map[string]TranslateSiteBlacklistWithTime{
				"1013309121859": {},
				"8181035976":    {},
			},
		},
		Media: Media{
			DeviceIDSalt: "5A63EA96076FE1C229C12EE8C2B9DD88",
			Engagement: Engagement{
				SchemaVersion: 5,
			},
		},
		MediaRouter: MediaRouter{
			ReceiverIDHashToken: "Gkb6cqJJ2D5bkFpwIWyFM8DUF9k93XyAUyCexMJxToGC/5EnIaowUBoHzadGtNsR8jnxm/2eHrZkCY531TMxSg==",
		},
		NTP: NTP{
			NumPersonalSuggestions: 10,
		},
		OptimizationGuide: OptimizationGuide{
			Hintsfetcher: Fetcher{
				LastFetchAttempt: "13337883014209791",
			},
			Predictionmodelfetcher: Fetcher{
				LastFetchAttempt: "13338424215008485",
			},
			PreviouslyRegisteredOptimizationTypes: PreviouslyRegisteredOptimizationTypes{
				AboutThisSite:   true,
				HistoryClusters: true,
			},
			StoreFilePathsToDelete: TranslateSiteBlacklistWithTime{},
		},
		Prefetch: Prefetch{
			SearchPrefetch: SearchPrefetch{
				Cache: Cache{
					HTTPSWWWGoogleCOMSearchQCreepjsOqCreepjsAqsChrome69I57J0I19I512L2J0I10I19I512L2J0I19I512J46I19I512J0I19I512J46I10I19I512L2991J0J7SourceidChromeIeUTF8: []string{
						"https://www.google.com/search?q=creepjs&oq=creepj&aqs=chrome.1.69i57j0i19i512l3j0i10i19i512l2j0i19i512j46i19i512j0i19i512j46i10i19i512&pf=cs&sourceid=chrome&ie=UTF-8",
						"13337882908802101",
					},
				},
			},
		},
		PrivacySandbox: PrivacySandbox{
			AntiAbuseInitialized: true,
		},
		Profile: Profile{
			AvatarIndex: 26,
			ContentSettings: ContentSettings{
				EnableQuietPermissionUIEnablingMethod: EnableQuietPermissionUIEnablingMethod{
					Notifications: 1,
				},
				Exceptions: Exceptions{
					ClientHints: ClientHints{
						HTTPSWWWGoogleCOM443: HTTPSWWWGoogleCOM443{
							LastModified: "13337882802178698",
							Setting: HTTPSWWWGoogleCOM443_Setting{
								ClientHints: []int64{
									9,
									10,
									11,
									13,
									14,
									16,
									23,
									25,
								},
							},
						},
					},
					MediaEngagement: map[string]MediaEngagement{
						"https://abrahamjuliot.github.io:443,*": MediaEngagement{
							Expiration:   "13345659587492709",
							LastModified: "13337883587492715",
							Setting: MediaEngagementSetting{
								HasHighScore:          false,
								LastMediaPlaybackTime: 0,
								MediaPlaybacks:        0,
								Visits:                1,
							},
						},
						"https://browserleaks.com:443,*": MediaEngagement{
							Expiration:   "13346200059144075",
							LastModified: "13338424059144078",
							Setting: MediaEngagementSetting{
								HasHighScore:          false,
								LastMediaPlaybackTime: 0.0,
								MediaPlaybacks:        0,
								Visits:                1,
							},
						},
						"https://discord.com:443,*": MediaEngagement{
							Expiration:   "13346200242282831",
							LastModified: "13338424242282835",
							Setting: MediaEngagementSetting{
								HasHighScore:          false,
								LastMediaPlaybackTime: 0.0,
								MediaPlaybacks:        0,
								Visits:                7,
							},
						},
						"https://www.google.com:443,*": MediaEngagement{
							Expiration:   "13345658804538028",
							LastModified: "13337882804538031",
							Setting: MediaEngagementSetting{
								HasHighScore:          false,
								LastMediaPlaybackTime: 0,
								MediaPlaybacks:        0,
								Visits:                1,
							},
						},
					},
					SiteEngagement: map[string]SiteEngagement{
						"https://abrahamjuliot.github.io:443,*": SiteEngagement{
							LastModified: "13338423195338550",
							Setting: SiteEngagementSetting{
								LastEngagementTime:     1.3338394395338456e+16,
								LastShortcutLaunchTime: 0.0,
								PointsAddedToday:       2.7,
								RawScore:               2.6664000000000003,
							},
						},
						"https://browserleaks.com:443,*": SiteEngagement{
							LastModified: "13338424058498129",
							Setting: SiteEngagementSetting{
								LastEngagementTime:     1.3338424058498124e+16,
								LastShortcutLaunchTime: 0.0,
								PointsAddedToday:       4.5,
								RawScore:               4.5,
							},
						},
						"https://discord.com:443,*": SiteEngagement{
							LastModified: "13338424236855637",
							Setting: SiteEngagementSetting{
								LastEngagementTime:     1.3338424236855626e+16,
								LastShortcutLaunchTime: 0.0,
								PointsAddedToday:       15.0,
								RawScore:               14.6813503758336,
							},
						},
						"https://www.google.com:443,*": SiteEngagement{
							LastModified: "13338423195338533",
							Setting: SiteEngagementSetting{
								LastEngagementTime:     1.3338394395338456e+16,
								LastShortcutLaunchTime: 0.0,
								PointsAddedToday:       3.0,
								RawScore:               3.0,
							},
						},
					},
				},
				PrefVersion: 1,
			},
			CreatedByVersion:                       "103.0.5060.53",
			CreationTime:                           "13302807096310804",
			ExitType:                               "Normal",
			LastEngagementTime:                     "13338424236855626",
			LastTimeObsoleteHTTPCredentialsRemoved: float64(time.Now().Unix()),
			LastTimePasswordStoreMetricsReported:   float64(time.Now().Unix()),
			ManagedUserID:                          "",
			Name:                                   generateRandomMD5(),
			WereOldGoogleLoginsRemoved:             true,
		},
		Safebrowsing: Safebrowsing{
			MetricsLastLogTime: "13338423195",
		},
		SegmentationPlatform: SegmentationPlatform{
			ClientResultPrefs: "ClAKDXNob3BwaW5nX3VzZXISPwo0DQAAAAAQgYmt1eHc2BcaJAocChoNAAAAPxIMU2hvcHBpbmdVc2VyGgVPdGhlchIEEAcYBBDUxdut4dzYFwpXChFjcm9zc19kZXZpY2VfdXNlchJCCjcNAACAPxCaia3V4dzYFxonCh8KHQ0AAAA/Eg9Dcm9zc0RldmljZVVzZXIaBU90aGVyEgQQBxgEENbG263h3NgX",
			DeviceSwitcherUtil: DeviceSwitcherUtil{
				Result: Result{
					Labels: []string{
						"NotSynced",
					},
				},
			},
			LastDBCompactionTime: "13338259199000000",
			SegmentationResult: SegmentationResult{
				CrossDeviceUser: User{
					InUse:         false,
					SegmentID:     1001,
					SegmentRank:   1,
					SelectionTime: "13338064251708246",
				},
				ShoppingUser: User{
					InUse:         false,
					SegmentID:     1001,
					SegmentRank:   1,
					SelectionTime: "13338064251708116",
				},
			},
		},
		Sessions: Sessions{
			SessionDataStatus: 5,
		},
		Settings: FingerprintSettings{
			A11Y: A11Y{
				ApplyPageColorsOnlyOnIncreasedContrast: true,
			},
		},
		Signin: Signin{
			Allowed: false,
		},
		Spellcheck: Spellcheck{
			Dictionary:         "",
			UseSpellingService: false,
		},
		SupervisedUser: SupervisedUser{
			Metrics: Metrics{
				DayID: 154378,
			},
		},
		Sync: Sync{
			Requested: false,
		},
		UnifiedConsent: UnifiedConsent{
			MigrationState: 10,
		},
		Updateclientdata: Updateclientdata{
			Apps: UpdateclientdataApps{
				Ihcjicgdanjaechkgeegckofjjedodee: Ihcjicgdanjaechkgeegckofjjedodee{
					Cohort:      "1::",
					Cohortname:  "",
					Dlrc:        6086,
					Installdate: 6085,
					Pf:          uuid.NewString(),
				},
			},
		},
		WebApps: WebApps{
			DidMigrateDefaultChromeApps: []string{
				"MigrateDefaultChromeAppToWebAppsGSuite",
				"MigrateDefaultChromeAppToWebAppsNonGSuite",
			},
			LastPreinstallSynchronizeVersion:  "113",
			SystemWebAppFailureCount:          0,
			SystemWebAppLastAttemptedLanguage: "fr-FR",
			SystemWebAppLastAttemptedUpdate:   "103.0.5060.53",
			SystemWebAppLastInstalledLanguage: "fr-FR",
			SystemWebAppLastUpdate:            "103.0.5060.53",
		},
		Webauthn: Webauthn{
			Touchid: Touchid{
				MetadataSecret: "FAs08eDqvux1A4NYorVc4ZHDwnhqyLknX9ef3JS4DLg=",
			},
		},
		Zerosuggest: Zerosuggest{
			Cachedresults: ")]}'\n[\"\",[\"rafael nadal\",\"joao felix fc barcelone\",\"tapis rouge mostra de venise\",\"sepp kuss\",\"luis enrique psg\",\"radeon rx 7800 xt\",\"incendie marseille entrepot\",\"rc strasbourg\"],[\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\"],[],{\"google:clientdata\":{\"bpc\":false,\"tlw\":false},\"google:groupsinfo\":\"ChwIkk4SFwoVUmVjaGVyY2hlcyBwb3B1bGFpcmVz\",\"google:suggestdetail\":[{\"google:entityinfo\":\"CgkvbS8wNTFxMzkSGUpvdWV1ciBkZSB0ZW5uaXMgZXNwYWdub2wy7w5kYXRhOmltYWdlL2pwZWc7YmFzZTY0LC85ai80QUFRU2taSlJnQUJBUUFBQVFBQkFBRC8yd0NFQUFrR0J3Z0hCZ2tJQndnS0Nna0xEUllQRFF3TURSc1VGUkFXSUIwaUlpQWRIeDhrS0RRc0pDWXhKeDhmTFQwdE1UVTNPam82SXlzL1JEODRRelE1T2pjQkNnb0tEUXdOR2c4UEdqY2xIeVUzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM04vL0FBQkVJQUVBQVFBTUJJZ0FDRVFFREVRSC94QUFiQUFBQ0F3RUJBUUFBQUFBQUFBQUFBQUFGQmdNRUJ3SUJBUC9FQURBUUFBSUJBd0lGQXdNQ0J3RUFBQUFBQUFFQ0F3QUVFUVVTQmhNaE1VRWlVV0Z4Z2FFVk1pTXpRbEt4MGZBVS84UUFHUUVBQWdNQkFBQUFBQUFBQUFBQUFBQUFBd1VCQWdRQS84UUFJUkVBQWdNQUFnSUNBd0FBQUFBQUFBQUFBUUlBQXhFU0lRUXhJa0ZSWWNILzJnQU1Bd0VBQWhFREVRQS9BREdpNnRLR3ZvcFRoV2NrSDROSm5FbXZRL3FFRnRBY2lOd3pueDNxM2M2azJuMnAzbzI0L3VJNjBvYWxLbDFLSlVYYUQ1N1ZkQjhpUkp1Q3JnQjJhc21xUlM2WXJLNmdoZTVvRGQ2eExkMnJXMGVwN2M5c0xrSDR4M29WUGVnaGRQc2NNeDlCSUhUSHhVa1BDbXJwR2Y4QXpRTHRCNitvWm9GbHU5UXRkTThuRjFGQVJic2toQTlXMG5PUGZCcjdnclcwMC9pRldWTnp5RGFjZUtzUjZOclVSWS9wMGo3UFZrUGcvYXF0cGJSWE9zV3N0dXZMSmwvaURPTnA4OVBIYnQ3MU5WZzNKMXFOeDdtd25XSG4ybFJ0WS9OV1UxSzZXTW1VQUR4UTRXTUloajNTYld3S2cxU0dTRzF4RklXeldrK3BtNVRNdFJlUVdEUGNxRG52Z2RxV1lvRnVFMnFjRFBmMnJidEowTzJ2TjViREsxSWZIZWdKb04yWklla1V2WUFZd2E1RUNlanNwYnJPRGtWdUc1YmlIakRseUtXQ0Z0b0k4WTZmakZhSXV0MjRrNWN0L0NySHNxU0Evd0NLV05OaHR4RkZkN1ZhV2VMbCtqb1ZKNk1EODV6WGsvQzJrcThidzNCNWpNTnE3em5QdDdVcHRJWnU0K3FyeGV1eEhJOFEyTm9GTXQraXNlZzNOZ1VwM2o3T01aYnEyYU5yU2NSeW95SDBzVDBPUHVEVjNVOUgwU2VheXVycWFNRUlJajVYZDE3ajVydlZZN0tHMHQwc1dSbGdrWGFGOXZVVCtUK2F2NDV4d0puODRjS1dmOFJ0dWJ3Y2lONUNjQUNwN2VXYlVJbWEyQlpFcEx2TmY1bG9JeEdjcVBJcnpobmphVFN5OGNrVytKejQ3aW1HTVcvVVUxM1Z2WDhmY3UyUEhFMmpwREE5akk2WTZzckNyZkV2RlZscldqR0piT1JwcE9paGtIUW1sV1c3dHhLeVM0M0FrZGE2czcyS2VSSVlNRmczWDdWVWp2VENjc1hvUzEraDNsbHcvY1RUQlNpam1MdEhxVDMvQU8rS1hiZU5weXpQTXZMY2JYM2sra2ZHS2JOZTR1dDlQUnROdVltZUtTQW84aW45akVlUjVwQUY1UHAwclMycFNlM2ZzRDFCK2xZN2xMSGxHbmlXaFZLeGdHbDI2MmpKWTNjVW5tVGFwUFFlUDNVWWowdVMxMEsxdW1CQ3lrc1VidUZKNmZqRkFOSTFHNjFUTVhLaHQ0bUlCRWE5WHozeWFjcE5adDlWdERHR0hLSktaUDhBU1FjRWZVRVYxUzlrbVI1amMxQ2ZSZzZTMFdXemJsUkZpUjNwTm1TVzBsTWM2Rkd6a0Exb0ZyY1Iyc1pnam1VNHBWNHJZM1U2T29YMCsxYTY3RHl3eGMxRk5hQVZMbjlrZHZhcWxqTnFHb3VUSTdrUXdLd1hkN2tud0s0ZzFrV2dLMnR2RERudTRiZitUL3FvTmZ2ZHVvelJLY0pFeFFEMnhRN2ZISU9vaUJQa2pyOTZKZ2drMEtKN2VRdmVzekNVU00zVTd1LzF5T2hvT2kzTmhLUW5yako5VVQ5alJPUnR2OGlXTlpWNnI2dlAzODFMelZ1NHdaRVV1UmhsWlIzSGNFVkRLQ01Nc3JGVG9oWFRMNjFqdGhNakNJbFNkaDdnanZWZlQ3MHRiU3BGT0ZrTE5JQnVIUTU4ME9oZ2lEcThTaU1BN1hHQVFBZk9LN2pkb3V5NEtuREFlL2FoMVVpc2tpR3V2YTBBSDZoRkpvbElhU0RsU2tlc283REorblVVeDZHbW1YMDhTVFNFT1R0Q3Y1UGpyU2k0a2xWT1ZGSTIxU2VpazRVZCszdFhFRSswQmszWlA5dUFjajIrYXV5QXdhdVZuLy9aOgxSYWZhZWwgTmFkYWxKByMyZjRmNzVSNmdzX3NzcD1lSnpqNHRUUDFUY3dOU3cwdGpSZzlPSXBTa3hMVE0xUnlFdE1TY3dCQUZhcEIyUXAGcAc\\u003d\",\"zl\":10002},{\"zl\":10002},{\"zl\":10002},{\"google:entityinfo\":\"Cg0vZy8xMWM3aGdnbHhrEhNDeWNsaXN0ZSBhbcOpcmljYWluMqsNZGF0YTppbWFnZS9qcGVnO2Jhc2U2NCwvOWovNEFBUVNrWkpSZ0FCQVFBQUFRQUJBQUQvMndDRUFBa0dCd2dIQmdrSUJ3Z0tDZ2tMRFJZUERRd01EUnNVRlJBV0lCMGlJaUFkSHg4a0tEUXNKQ1l4Sng4ZkxUMHRNVFUzT2pvNkl5cy9SRDg0UXpRNU9qY0JDZ29LRFF3TkdnOFBHamNsSHlVM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOLy9BQUJFSUFFQUFRQU1CSWdBQ0VRRURFUUgveEFBYkFBQUJCUUVCQUFBQUFBQUFBQUFBQUFBR0FnTUVCUWNCQVAvRUFEUVFBQUlCQXdJREJRVUhCUUFBQUFBQUFBRUNBd0FFRVFVU0V5RXhCa0ZSWVhFaU1wR2hzUWNVRldLQndmRVdJMEpTa3YvRUFCa0JBQUlEQVFBQUFBQUFBQUFBQUFBQUFBSURBUVFGQVAvRUFDWVJBQUlDQVFJRENRQUFBQUFBQUFBQUFBRUNBQkVESVRFU1FYRUVCUlFpVVlHUnNkSC8yZ0FNQXdFQUFoRURFUUEvQU1reFJGMmIwQmRRQ3p6cVhWbTJ4eGc0M0h4UGxtaDdGSCtrWFVGaHBkbVpwMGhZUkRJWndyY3h6Nit0UkpsaW1sUGJwSWlRb2doSE5Vd2U4REF4Ni9JMGg3Q1ZpNmdJU2lLN0FzQmhXWElKejNjd1BVZ2Q0cGxlMFZqeEN5WDhJa09NdHhqaytQUFBlYWZtMUUza1FDeWlTTUxnQlpTUjhNMXd2bkQ4cE9rcko0Z1VLak8wOVJRZnI4UERNZURrZU5GODdrSEE1VU9hMGdrc25ZOVZiY0RYU0RCdzExYThhNm9yb01zVWxlRkR3dVRuL01kY2VYaFROdHA5NXFVckZVZkFiQko4YWQ2VVNhQnFFbGhaelRwYkpNaGtVRUVuSU9PZ0FCNjR6UU1TQkR4b3JOWmtTMjdFWERxTjVBSkhoVWU4N0tYOWdUSkNXVS83SWNHdEMwUFg0OVZESW1uend5SW0vRHBqSThxcDd2WHRYdjUvdTl2cFRSUkU0M2hONVA4QTF0RklEdGN2SEZqcmFCb3ZkWHNobVhNMFk2OFFmdlhwZFJpdmRPbjJLVWtYM2tQcjFGRzM0YkpjV01nMUJDVWNZRzVBckEvb2NVTGRudENrZTlsNHBCaUJkR0RJQ0R6eDM5UHJSak5vYmltN01iQVhuQm8xMEN1eWJPSS9DSk1lNDdDZThaNWZLdkNueW5KK0tOUHMwbGpNOS9iVG4yR2pWeHp4MEpCK29vTXF3MEsvL0RkU1NkaG1NZ3BJUHluK0FmMG9YRmlvekUzQzRNMXBMaXhnUnpHMFVhN0dBM05ndDRta3czRmxORnVDSlBFcDVTS3VSK25qVUJZYmU1dG8yVkROQzQzS2l4aVFEMDhxazdOcFZwSG5Dcnoyc3FyOUJWUXJOWlQ2Um5WSFJ4aFBkb1IxalgwczdLODB5M3RuUzVDazhVNDI0YlBNZWZXaVcvdUM1SEk4K2FxUHBRVjJodHlzOTNMSjc3S3FmQVorcG9rQUxheFdkbUNXSUlVdFJYRFQ4TUpZamQ3STg2dHpLQXVTY1Y3dnBRRktTS1NWMWpnamFTVmpoVVhxeDhCWFNZVTlodFJ1b2pQYUt4NEdONkU5RmJ2QTlSei9BSnEvbTFTYVZtaTJzVzZFa1ZIZnNYcnVpd2FmT24zV2VSMEFBTE9oUnlNRVpBSVBNbnFCOGFuMlhaUHRUcXNNTW5GMDYxZ2xSWFdVR1NSc0hCeUZLcjh6U0dUSVNkSm9wa3dvcWpqNjZiUkZ2SkhFcnl5WmVRTGtudVVWU1hFTVdvOFJKZDl3MjcyMXRVYVVvZkE3UWNZODYwU3crem15aDRiYXBQTnFaSE5vNS9aaEhwRW1BZS9rYzkxVzZXdDBWNE9uMjBlazJFUTUzTXlLRzI5NGpqNktQTnZoU3ZERzdadmlYRTd3eDhKVEVnbzdsdndmUXVZSmM5a05XRndHc05LMU9lQSs2WHRpRG45NnZ0SCt5N1hMNkY1dFFUN2p5L3RSc2drWW44MkQ3STZlZmxXNk5mMnphYmMzVmhKSGVpM2pZNGlrRDdpb3pqSTc2aXhyL1VmWlpRWmpBYnlBSzd4RDNUMFlEeXlDUFNyWXJiZVpKeHRxNUhDTHJwN2J6Ly9aOglTZXBwIEt1c3NKByM0MjQyNDJSN2dzX3NzcD1lSnpqNHRWUDF6YzBURGJQU0VfUHFjZzJZUFRpTEU0dEtGRElMaTB1QmdCc01naUFwBnAH\",\"zl\":10002},{\"zl\":10002},{\"zl\":10002},{\"zl\":10002},{\"google:entityinfo\":\"CgkvbS8wNG1wOHgSEENsdWIgZGUgZm9vdGJhbGwyxhdkYXRhOmltYWdlL3BuZztiYXNlNjQsaVZCT1J3MEtHZ29BQUFBTlNVaEVVZ0FBQUVBQUFBQkFDQU1BQUFDZHQ0SHNBQUFCUVZCTVZFWC8vLzhBbitQY0x6UUFudU1Bbk9JQWwrRUFtdUxRME5ETHk4djM5L2NBbmVYVDA5SzN0N2NBTG92czdPejQvZi9FeE1UYTdQbSszdlZacytnQUs0dmg4UHZvOC93QXBld0FQSkVBU1phSHhlMDZxZWJkM2QyajBmRW1wT1R0OS8zSTQvWnF1ZXF0MXZMZUlpaFpkNjBBTTQwQWpkTjd3T3l3dXNGS3J1ZmNLUzRBajk4QWc4WjNqcm1KcHJxT29NTUFRcE55ajZOR2FxYS95TnptSUJjQWU3ZGtsYldlc2I1QWg3SE1YMkFBYXBzRWdMZnNwNm5kRkJ3dVJHRlZpYXI1Nk9sUUlDWUFrYzNmVTFlSlFEand3Y0txVWt5M1NFZWpzTEc3S3kzQUNoUEpBQURHdXJqb21KcXhsNWJiUFVLbWRuZXFob2ZBcTZ5dVhGMnFqS3F2YVdwY2g3KzVIU0NhS2tZblY1eTZVR3FXYXBPYVQxRFQyZWVrc3MrZkhpSUFGWVRFTzFTQ2hvZ3lXM1NrUVVOZFYzdXBYb0M5UTEzb0FBRGlZMlo3ZXFxSjI5dkpBQUFIUGtsRVFWUlloYVZYQzNmYTJCRytRbGZYa2tBQ1NUd2xKSVFFS0NnR2JBTzJhNjlLbXQxa2QwT0tIYWVKNis3YVRyWnBtdWIvLzRETzZBVjJjRGJuWk02eEQwY3duMmErKzgzakV2S0FxVHRneGRqd2svclE3N1piV3hSVGozSTV4UlBGOXJkNlN4VlJJcVNqbXdZbnhNWVpwdDZCNTJKRitpYjNJaUUxMHhFWTVYa3VOcDZuVEhETUdpSEZQNGNRd1YyM0JBcStsS2VVb3h6SEdQN25xV0RwQUNGKzFWMnRnTHZEZUhEbWVOT3lBMWRudktieEpzVndlTVlCUk9VcmZBSlBkVXZBbDV1R3l4RmI5d3hkQ0V6ZWNvVWdHSEtNNXdXckRndy81RitTaUMxZ3RNeTB1MTNUdFR1bVhSTUN6d3hNd1dVYTFUa0lUYkNKVkhyQW54Q0RnVGVybVVQUDZibzJOWVd1QTJoZFV3czBUcmZOR256Tk1TUCs2VGIvc2dPa0dUclZQVnZYTEFlNTR4TXlHUis0bHVaNmpPL0NVNmU4RGFGRU9qdzRXSFV6c0hVVDgrVTJqVExMclEwRlZ6ZmhKWHpuU3dSNFAvakRtZW11WmpzMGRwSmxtVEVHLzJNc3lJMmF0cVVQUVJiOEZ6R0E5bEl2cHNVaTRIaDVmbmhVZ1dxb1BCb2Y3eVlZMU5acmxxYkJCd2QwdWVtdnRzbGZUaHBwc0luNlRqZGZJWldPQndoQkhXQmlhTUZoR0tTOXFZY0srU0dLL3JxUk1uOTZYN1B0aFlEUEhXMElNSlJqTmppdEV5QlBab1dDLzdlbmpReEEza0p6NVZSR0JFTXdOWnR5UXAza1NVaEY4bU5ZQUlTZmp1VTBBS0g0SlFDUkZqTHEyUnZhWFRnS2l4U3pLQ3ZrV1ZSQTg1L3Zzd1NCN1d3QklHUWZ5OFRRN01jYTR3UTlTd0lDT0FzTENjTFB2enlOK1phUHRnS1FmUmtsSVhpb0tTY0xJUThBTFB5MUVxZkJuMjRISUF0RTBMdVBMZEI0Rm9LWU1KQWl2TmdaRHpDSGpFWHhhSHg0T0Q3S09vbDBpbDhPWGRlanZKWFFXSlNlNUFHZ0xWL3V6Q0dOZWZ6ei9Ua3FFU1c1dTBoZk44QnF0ejFrb1NhMTR3RCs3bThDRktLVmRDeno4aGkrTzVFM3BDRWN4MGV6d1BwQ05RQk9ISUtZVTVnam5MZC9PVzNzcXFUQzdoUVVQMEJxMjROWTA1NU5nVVlBVU5XN0dTUUlSSjAzRHNrak9Tc3FzUGgwOTlNUUJOME5RRXdkVlNVNzVObEdCcUh2K3hDUC82cmMzdDFWRVlCbkovdEhwZExSL3NrQVZRU0pWZkFnYWc3cVdRZjNuVTBLb3Vlcmk5WHJGMUhvdjdob2o4bVJ6QW1MWExEaUFpQUdGU0x0WWc2RzVsRWdBUUgrTWN2OUwrSmZsa3R2WGtTZlZvU001WG1KU0IxSjZxU1NuVGU0dVVTT29hKzVYaGNBakJpZzhUWktXQXhmNTNvcFYxNjFpYm83YjVPcVVnVlRSdkZqOWFRQkdoM0xVRk1tYzZCcGdYdVJDQTE2Rm1maFgyeHFUcjI0L0dlYlNNclZDS3luOUpLSDg4WWN1YVdCMXVtQUVzQWRBSURlZitGSlJPc0thcis1WFBwUnFRd0FpZWRCTmMxaXdDb2lSTkMxYXhwTkFNb0kwUGdOaUlqeStWdGNRa2poNytRNkJwQWtTVzBwaEV5dURxNmhHQTZMRE1Wb1dRQlF6Z0E0SkdJTmNJNHBSU0s1UVlCSnRkOVhsQWtaOVp2OWFxYzRPQVVBNkV5dW13TEVLUURDMDdOWjNrVitSbFp2NFowSWNLMjA5bnBUUXBRV3VYNTNRNDRITHpFRmozTzVEUTVpZS84eUE3aEVTcytKMUx3aGtFTDlxb2xFQUFDWnFtVE0va0FwQmw2UWNMQkR1R3lFckx2SVpYSW05ZXFJeENTT3FxMDZwb0NCbE9SL056ZyswTENndVJqQW9GbXhMRFpUaUVya3VwOENrR216UHlXVFBVVzVBaVdmUUZHNEhBU2VDc25NQUZCbGliMU9BQ2JOSy9JT0FXNWFlM3RWZUh1OTE3OFIyU2syWjEwek15bnJlZEhLV2F0ZitYRUtDRkFEMEY2LzN4dnRIZUEzU3E4U3R6eE43NmJGcEtxZG5NVThCekhDaW93Qk1QNytGYUtnbE9CUWpyQVl1NTZuT1h4Y3p0QVRuSHdRRHpJbEFJdmhMUUxjVENaU1Q3bEdHR1YwZlgyZ1RJK3hJdzBkVFdkSlE0Ry9OUWw1Q0s5OEZPYWsyV28ycS9VREpWR3pvaWpWa1lyRmJHamFrR1V0clMzVjhodzRJV1ZCWFVJSWI2YjlWcXZWbkk2VVpJNU9SNU1wMmNjQWJOdGFOMVZBc2RiTFJEWlVWOERDc3ZpMjJXcnRUYWJWVVY1bHhkMDRVcysyODdZTzB5RStCMXhrNENDeUpENkVoZkJTNmpYMzl2cmxYclUzUlFNUy81UE1YMHJqTTBqaWhkRUdORUtMY0pGTmVUOEI2RndDd20xNWNuQndNQUlsb3IyYmxqL01QbVlUZkQzYU1BU0JHblhQeFJVeDZieEl3NjBQTWFUMVZjY0FydHUzWWNGL245R1ZEMWNNd2FJOEk0SWJJSUs4U0lETHYwZGh1RHpQYTF4OXRjUWlEUXZ4R3JFNTNvR0xPaHlFb1hWaGowS0UwN1FUVjI0ajMxOStXSlZFc2JSNnZreTdkemo3REFpYkN3WW1ZVFBlMGdRTGVuVWMzaUo5cjNqKzZ6S2FmZm8wOHk4djE5Tmo5bHZqN29xRFM1WUJUVG9ZNmdGbHlSUmJsTklBMVNLOC80K2Y3bzZ1Lzk1YnNwSTFqdzdkMm1QVGpkT0FOVy8zZVB5b0l1NklMMWV2TC8yNzR6TXMzRnZ6MGtXVFdqQnZiTXVtaWJCNEhPdnl4MC8zdk1FL2xMYXRxckRxOGtPalk5VUN6NkJyYlRZK3o4TDcvaysyTDh1d2JNTnViZ1lDYnBNYkNFOExkeEQ4c3kzdlR4QnczYWVPWnRwdU55L1FwTjF1RFBEb2Z3K3MrOW1GZzJkQ0xlRHVXdU5qTm9ERDJROFBYampXVng0aGNPNGpmSTRSd3VqSEoxKzU4cXd2WFR4MzN4cHZ3ekNNQ3MrK2Z1a2lkNjU5OTR5K1AzdjJwOWMrOHYwWFQvTGRWOS9ZdnUveW5YcDg0L1gvLzNaL3kxa2NHTCtyQUFBQUFFbEZUa1N1UW1DQzoNUkMgU3RyYXNib3VyZ0oHIzAyNDg5NlI3Z3Nfc3NwPWVKemo0dFRQMVRjd3lTMndxREJnOU9JdFNsWW9MaWxLTEU3S0x5MUtCd0JvOEFpWnAH\",\"zl\":10002}],\"google:suggestrelevance\":[1257,1256,1255,1254,1253,1252,1251,1250],\"google:suggestsubtypes\":[[3,143,362],[3,143,362],[3,143,362],[3,143,362],[3,143,362],[3,143,362],[3,143,362],[3,143,362]],\"google:suggesttype\":[\"ENTITY\",\"QUERY\",\"QUERY\",\"ENTITY\",\"QUERY\",\"QUERY\",\"QUERY\",\"ENTITY\"]}]",
		},
	})

	return data, err
}

func Setup() (string, error) {
	randomMD5 := generateRandomMD5()

	sourceDir := "C:\\Users\\arm\\Desktop\\MYBROWSER\\gologin\\base"
	destinationDir := fmt.Sprintf("C:\\Users\\arm\\Desktop\\MYBROWSER\\gologin\\prof\\%s", randomMD5)

	err := copyDir(sourceDir, destinationDir)
	if err != nil {
		fmt.Println("Error copying directory:", err)
		return "", err
	}

	defaultDir := filepath.Join(destinationDir, "Default")
	err = os.Mkdir(defaultDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating 'Default' directory:", err)
		return "", err
	}

	jsondata, err := GetFingerprint()
	if err != nil {
		return "", err
	}

	preferencesFile := filepath.Join(defaultDir, "Preferences")
	err = writeJSONToFile(jsondata, preferencesFile)
	if err != nil {
		fmt.Println("Error writing JSON data to Preferences file:", err)
		return "", err
	}

	fontsDir := filepath.Join(destinationDir, "fonts")
	err = removeRandomFiles(fontsDir, 64)
	if err != nil {
		fmt.Println("Error removing random files from 'fonts' directory:", err)
		return "", err
	}

	fmt.Println("Operations completed successfully.")
	return destinationDir, nil
}

func generateRandomMD5() string {
	h := md5.New()
	_, _ = io.WriteString(h, randStringBytes(32))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func copyDir(src, dest string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, info.Mode())
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func writeJSONToFile(data []byte, filePath string) error {
	return os.WriteFile(filePath, data, 0644)
}

func removeRandomFiles(dirPath string, numFiles int) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	if numFiles > len(entries) {
		return fmt.Errorf("not enough files in the directory to remove")
	}

	rand.Shuffle(len(entries), func(i, j int) {
		entries[i], entries[j] = entries[j], entries[i]
	})

	for i := 0; i < numFiles; i++ {
		fileToRemove := filepath.Join(dirPath, entries[i].Name())
		err := os.Remove(fileToRemove)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateRandomIP() string {
	octet1 := rand.Intn(256)
	octet2 := rand.Intn(256)
	octet3 := rand.Intn(256)
	octet4 := rand.Intn(256)

	ipAddress := fmt.Sprintf("%d.%d.%d.%d", octet1, octet2, octet3, octet4)

	return ipAddress
}

func randomElement(slice []string) string {
	index := rand.Intn(len(slice))
	return slice[index]
}

func randomElementInt(slice []int64) int64 {
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

	renderer := randomElement(gpuSlice)

	return el, renderer
}

func randomHex(length int) string {
	bytes := make([]byte, length/2)
	rand.Read(bytes)

	hexString := hex.EncodeToString(bytes)

	return hexString
}
