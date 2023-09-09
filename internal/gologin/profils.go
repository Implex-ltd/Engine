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
	"strings"
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
			PrevNavigationTime: "13338531332965699",
		},
		AccountIdMigrationState:         2,
		AccountTrackerServiceLastUpdate: "13338531333004494",
		ACKExistingNTPExtensions:        true,
		AlternateErrorPages: AlternateErrorPages{
			Backup: true,
		},
		AnnouncementNotificationServiceFirstRunTime: "13338531332855366",
		Apps: FingerprintApps{
			ShortcutsArch:    "",
			ShortcutsVersion: 0,
		},
		Autocomplete: Autocomplete{
			RetentionPolicyLastVersion: 115,
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
			EditingEnabled: false,
		},
		Browser: Browser{
			EnableSpellchecking: false,
			HasSeenWelcomePage:  false,
			WindowPlacement: WindowPlacement{
				Bottom:         1310,
				Left:           758,
				Maximized:      false,
				Right:          2678,
				Top:            230,
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
		CredentialsEnableService: false,
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
		DipsTimerLastUpdate: "13338531332959012",
		DomainDiversity: DomainDiversity{
			LastReportingTimestamp: "13338531333004017",
		},
		Extensions: Extensions{
			Alerts: Alerts{
				Initialized: true,
			},
			ChromeURLOverrides: TranslateSiteBlacklistWithTime{},
			Commands: map[string]Command{
				"windows:Ctrl+Shift+F": {
					CommandName: "humanTyping",
					Extension:   "mnldjpplkbaacjnffmcjmedffbaekgko",
					Global:      false,
				},
				"windows:Ctrl+Shift+Q": {
					CommandName: "automationTask",

					Extension: "mnldjpplkbaacjnffmcjmedffbaekgko",
					Global:    false,
				},
			},
			LastChromeVersion: "115.0.5790.82",
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
			ClientRectsNoiseEnable: false, // fuck
			DeviceMemory:           8192,
			DNS:                    "",
			DoNotTrack:             false,
			GeoLocation: GeoLocation{
				Accuracy:  100,
				Latitude:  49.4429,
				Longitude: 1.1003,
				Mode:      "prompt",
			},
			GetClientRectsNoice: rand.Float64(),
			GetClientRectsNoise: rand.Float64(),
			HardwareConcurrency: randomElementInt([]int64{4, 6, 8, 12, 16, 32, 64}),
			Icon: Icon{
				Avatar: Avatar{
					Enabled:  true,
					FullSize: true,
				},
				Text: "",
			},
			IsM1:       false,
			LangHeader: "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7",
			Languages:  "fr-FR,fr,en-US,en",
			MediaDevices: MediaDevices{
				AudioInputs:  randomElementInt([]int64{0, 1, 2, 3}),
				AudioOutputs: randomElementInt([]int64{0, 1, 2, 3}),
				Enable:       true,
				Uid:          randomHex(58),
				VideoInputs:  randomElementInt([]int64{0, 1, 2, 3, 4}),
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
				AllEnable:   false,
				FlashEnable: true,
			},
			ProfileID: randomHex(24),
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
			UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
			WebGl: WebGl{
				Mode:     true,
				Renderer: renderer,
				Vendor:   vendor,
			},
			WebRTC: WebRTC{
				FillBasedOnIP:  true,
				LocalIPMasking: true,
				Mode:           "public",
				PublicIP:       "",
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
						Value: map[string]int64{"0": 1, "1": 1024},
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
						Value: map[string]int{"0": rand.Intn(100000) + 1, "1": rand.Intn(100000) + 1},
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
			SavingDisabled: true,
		},
		Intl: Intl{
			AcceptLanguages:   "en-US,en",
			SelectedLanguages: "en-US,en",
		},
		Invalidation: Invalidation{
			PerSenderTopicsToHandler: map[string]TranslateSiteBlacklistWithTime{
				"1013309121859": {},
				"8181035976":    {},
			},
		},
		Media: Media{
			DeviceIDSalt: strings.ToUpper(randStringBytes(32)),
			Engagement: Engagement{
				SchemaVersion: 5,
			},
		},
		MediaRouter: MediaRouter{
			ReceiverIDHashToken: "Gk7tpVqhLJPezOj91Ls8wsRjia/sqLrzVD9BJwhV2BXYYx+TOsbyA4McT+pdBxOXiDfwMVVXak0x/1g8tdWYcg==",
		},
		NTP: NTP{
			NumPersonalSuggestions: 10,
		},
		OptimizationGuide: OptimizationGuide{
			Hintsfetcher: Fetcher{
				LastFetchAttempt: "13338531345761512",
			},
			Predictionmodelfetcher: Fetcher{
				LastFetchAttempt: "13338531345761512",
			},
			PreviouslyRegisteredOptimizationTypes: PreviouslyRegisteredOptimizationTypes{
				AboutThisSite:   true,
				HistoryClusters: true,
			},
			StoreFilePathsToDelete: TranslateSiteBlacklistWithTime{},
		},
		/*Prefetch: Prefetch{
			SearchPrefetch: SearchPrefetch{
				Cache: Cache{
					HTTPSWWWGoogleCOMSearchQCreepjsOqCreepjsAqsChrome69I57J0I19I512L2J0I10I19I512L2J0I19I512J46I19I512J0I19I512J46I10I19I512L2991J0J7SourceidChromeIeUTF8: []string{
						"https://www.google.com/search?q=creepjs&oq=creepj&aqs=chrome.1.69i57j0i19i512l3j0i10i19i512l2j0i19i512j46i19i512j0i19i512j46i10i19i512&pf=cs&sourceid=chrome&ie=UTF-8",
						"13337882908802101",
					},
				},
			},
		},*/
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
					/*ClientHints: ClientHints{
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
					},*/
					/*MediaEngagement: map[string]MediaEngagement{
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
					},*/
				},
				PrefVersion: 1,
			},
			CreatedByVersion:                       "103.0.5060.53",
			CreationTime:                           "13302807096310804",
			ExitType:                               "Crashed",
			LastEngagementTime:                     "13302807261018135",
			LastTimeObsoleteHTTPCredentialsRemoved: float64(time.Now().Unix()),
			LastTimePasswordStoreMetricsReported:   float64(time.Now().Unix()),
			ManagedUserID:                          "",
			Name:                                   generateRandomMD5(),
			WereOldGoogleLoginsRemoved:             true,
		},
		Safebrowsing: Safebrowsing{
			MetricsLastLogTime: "13338531332",
		},
		SegmentationPlatform: SegmentationPlatform{
			//ClientResultPrefs: "ClAKDXNob3BwaW5nX3VzZXISPwo0DQAAAAAQgYmt1eHc2BcaJAocChoNAAAAPxIMU2hvcHBpbmdVc2VyGgVPdGhlchIEEAcYBBDUxdut4dzYFwpXChFjcm9zc19kZXZpY2VfdXNlchJCCjcNAACAPxCaia3V4dzYFxonCh8KHQ0AAAA/Eg9Dcm9zc0RldmljZVVzZXIaBU90aGVyEgQQBxgEENbG263h3NgX",
			DeviceSwitcherUtil: DeviceSwitcherUtil{
				Result: Result{
					Labels: []string{
						"NotSynced",
					},
				},
			},
			LastDBCompactionTime: "13338431999000000",
			/*SegmentationResult: SegmentationResult{
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
			},*/
		},
		Sessions: Sessions{
			SessionDataStatus: 1,
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
				DayID: 154380,
			},
		},
		Sync: Sync{
			Requested: false,
		},
		UnifiedConsent: UnifiedConsent{
			MigrationState: 10,
		},
		/*Updateclientdata: Updateclientdata{
			Apps: UpdateclientdataApps{
				Ihcjicgdanjaechkgeegckofjjedodee: Ihcjicgdanjaechkgeegckofjjedodee{
					Cohort:      "1::",
					Cohortname:  "",
					Dlrc:        6086,
					Installdate: 6085,
					Pf:          uuid.NewString(),
				},
			},
		},*/
		WebApps: WebApps{
			DidMigrateDefaultChromeApps: []string{
				"MigrateDefaultChromeAppToWebAppsGSuite",
				"MigrateDefaultChromeAppToWebAppsNonGSuite",
			},
			LastPreinstallSynchronizeVersion:  "115",
			SystemWebAppFailureCount:          0,
			SystemWebAppLastAttemptedLanguage: "en-US",
			SystemWebAppLastAttemptedUpdate:   "103.0.5060.53",
			SystemWebAppLastInstalledLanguage: "en-US",
			SystemWebAppLastUpdate:            "103.0.5060.53",
		},
		Webauthn: Webauthn{
			Touchid: Touchid{
				MetadataSecret: "FAs08eDqvux1A4NYorVc4ZHDwnhqyLknX9ef3JS4DLg=",
			},
		},
		Zerosuggest: Zerosuggest{
			Cachedresults: ")]}'\n[\"\",[\"sergio ramos sevilla\",\"carlos alcaraz alexander zverev\",\"cyril hanouna tpmp\",\"accident police elancourt\",\"armée française au niger\",\"marché immobilier\",\"ligue 2\",\"mathieu kassovitz accident moto\"],[\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\"],[],{\"google:clientdata\":{\"bpc\":false,\"tlw\":false},\"google:groupsinfo\":\"ChwIkk4SFwoVUmVjaGVyY2hlcyBwb3B1bGFpcmVz\",\"google:suggestdetail\":[{\"google:entityinfo\":\"CgkvbS8wN2M0Z2QSM1NlcmdpbyBSYW1vcyDigJQgRm9vdGJhbGxldXIgaW50ZXJuYXRpb25hbCBlc3BhZ25vbDKDDWRhdGE6aW1hZ2UvanBlZztiYXNlNjQsLzlqLzRBQVFTa1pKUmdBQkFRQUFBUUFCQUFELzJ3Q0VBQWtHQndnSEJna0lCd2dLQ2drTERSWVBEUXdNRFJzVUZSQVdJQjBpSWlBZEh4OGtLRFFzSkNZeEp4OGZMVDB0TVRVM09qbzZJeXMvUkQ4NFF6UTVPamNCQ2dvS0RRd05HZzhQR2pjbEh5VTNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTi8vQUFCRUlBRUFBTmdNQklnQUNFUUVERVFIL3hBQWJBQUFCQlFFQkFBQUFBQUFBQUFBQUFBQUZBQUlEQkFZQkIvL0VBREVRQUFFREF3TUNBd1lHQXdBQUFBQUFBQUVDQXdRQUVTRUZFakVUUVFaUllRY1VJaktSc1VKRFVuR0J3U05pb2YvRUFCa0JBQUlEQVFBQUFBQUFBQUFBQUFBQUFBRUVBQUlEQmYvRUFDQVJBQU1BQWdJQ0F3RUFBQUFBQUFBQUFBQUJBZ01SSVVFU0V5SWpNUVQvMmdBTUF3RUFBaEVERVFBL0FQWkFtdTdhZUJYYlVBbVc4WStJcGVrOU9GbzBORXZVM2s3d0hEWnRsRjdiMStkeUNBTVhzYzRyRks4VytQZEdsKzhhakZnNm5DdmRiVENPbXRJLzFQUDF2Ui9YZyt2eExxRTJJMnRUalRUVWNKS1FkMjNjckEzRHV2ekZEOVpkZmxOdE5SMnJPN0FwNG5PeTk4VzNEdVBXbEx6VXE0SHNmODhPUGwrbm9PajZqRzFqUzQyb3dpb3NTRUJhZHdzUjVnanNRYmcrb3EyUldXOW1PeHZ3ejdva3FLNDBsNExCSEc1eFNoYno1clcycHFYdGJFNm54YlJFUlhLbHQ2VjJpVkhXcnRkdFNVUWtYVVFCNjFDSG5makdjclN2RktXR3lrR2ZIQzJ5dmdMU1NEOVJhczdJbnlZRGk1Y21SSGVhRE8xMUNVV3plL056OFdhTisyRFJwbXFOd0p1bk43bFJkd0pIelpJdDlxODNZMG5XOVRkUWg5cDFEWlh0Y0tra0FqOVhyU09XTld6cDRNbjFvOXY5bjBKVWJ3dkZlZENnOU1IdkxnVUxiU29ZSDB0V2p0UW5SSndSRGpRNVlEY2x0aEFVTzNGdWV4eHhSaW5aV2trYzY2YnB0aktWT3RTb2xSajdvYUF4ZFI0RlZWRXJPNWZOUXlKS0haTG9Rb0hwSzZaendiWC9BTHFwcXVyeDlNaTlkL2NxNXNoQ1BtV2ZTaUFJbEtYRzFJY0FLVGdnME1oNk0wMjhwMTQzVWxSNlFCd0IyUDcxZlpkUTgybHhwUVVoUXVsUTcwNmc1VGFiTEsybHBka1RVVnBwU2xJQnVya2szSnFadHhiUHlaVCtrLzFWUTZpMm1jSXVGS3h1TzRmRGZqSE5YckE4ZDZpcFBlZ05hTGFGSmNRRko0TktxR2pUNDAzM2xNUjRPQmx3cFZic2UvOEFHT2FWRUFMVEc2RDB0emNTSDNTdTIwQzEvdno5cWFtZWlPMEhGcEtrbDBOcEl0eVRqSklIT1BwV1M4VitLTlMwVFhOUVlEamJqQ0hMaERyVjlnS1FjRUVHMmZXaitoSzkvd0REOE42UXB0YXBDUThvc2toSUpWdUZqZStNZlNxVGFwNk5LeHVaMit3aXZYdE9RVnA2Nml0RzRGUFNXTWpkY1hJdCtGWGY4S3ZJMUl2V0lDVWJ5OHV4M2JmOFMvaXNRRGJHY2tjVUVjaG82YjZ6cHFsbENIMWpMdTVaK0lvL0ZtL1hjdjhBenhZV2NJb0w3YTF3Z215a28vTXdPcXJPVlc1MnEvWW10OUl5Mnd3MmpUWms4UE5PaGNobzNJUWZRYytuQjhzWHFMV05handIbFJIbWRSY1c2MWRJaHduWFRZM0h6SkJBT0R5UlUyalJHVXc0MGtSdWc1MDFCTFovTFNvZ2xQOEF4UDBvRDdRV25ZNkltb3R4bzBoQUpaV0gwM0NlNlNDUWJkKzNsa1ZtL0dTeVRyZ3ErejdVNUNKYzVic1RweGdTeXZGbkE0Z2pCU0xnWDNLUFBsU29acFR1b0JUcFdZN1ljVVZwTENWWFRnQXAzRW00eGZJNTR4aWxXUHQ1WnI2bjJqLy8yUT09OhRzZXJnaW8gcmFtb3Mgc2V2aWxsYUoHIzQyNDI0MlJBZ3Nfc3NwPWVKemo0dFRQMVRjd1R6WkpUekZnOUJJcFRpMUt6OHhYS0VyTXpTOVdLRTR0eTh6SlNRUUFxNDRMRWdwAnAGcAc\\u003d\",\"zl\":10002},{\"zl\":10002},{\"google:entityinfo\":\"CgovbS8waDk2cndmEhtDeXJpbCBIYW5vdW5hIOKAlCBBbmltYXRldXIygw1kYXRhOmltYWdlL2pwZWc7YmFzZTY0LC85ai80QUFRU2taSlJnQUJBUUFBQVFBQkFBRC8yd0NFQUFrR0J3Z0hCZ2tJQndnS0Nna0xEUllQRFF3TURSc1VGUkFXSUIwaUlpQWRIeDhrS0RRc0pDWXhKeDhmTFQwdE1UVTNPam82SXlzL1JEODRRelE1T2pjQkNnb0tEUXdOR2c4UEdqY2xIeVUzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM056YzNOemMzTnpjM04vL0FBQkVJQUVBQVFBTUJJZ0FDRVFFREVRSC94QUFhQUFBQ0F3RUJBQUFBQUFBQUFBQUFBQUFFQmdNRkJ3SUIvOFFBTXhBQUFnRURBd0VGQmdRSEFBQUFBQUFBQVFJREFBUVJCUkloTVFZVFFWRmhJakpDY1lHaEl5UnkwUlJEa2JIQjRmSC94QUFaQVFBREFRRUJBQUFBQUFBQUFBQUFBQUFCQWdVRUFBUC94QUFnRVFBQ0FnTUFBZ01CQUFBQUFBQUFBQUFCQWdBUkF4SWhNVUVpVVdFRS85b0FEQU1CQUFJUkF4RUFQd0ROaFZmZHpxWFlNY291UmdET0Q1MGJNL2R4TytDZHFrOFZTcXdkaG5jMkJ6UWpUcFlHNFl2N2JFQUhQblU1dHdUbG00WmNiaWZIeEorOUUyMW5mVEVDSzJsWk1kZGhPZmxVcTIwNnNFa2pkWFRnS3lZUDE5ZVRYV0lkVDlRZUdCNDdnNDJsUXZMNTRIci9BS29MZEgzM3NodG1mQ21WZXoycDNhZmhXRTBxRThqWnhuam1wcmpzcnFrRnE5dzlrRlZWeXlkVGlodW85d2pHNTdVcHJTVWs3RzU4YzV6UlZWaUh1THhRaWxBeEdWNjFaMFlzNGw1aWNaeDdKNW9iczdIQ2J1SXlybFF3SnpSWjVHRDBvVFQ3WjdlK2lUZCtHeHlQcFFieEhRZEUwelNKbVhDRDNRU0JqdzVwbmhnZ3VZL3pVRWNuNmxCTlp6WlhXb1J1eHRSQTBhQWJoSU92clRQMmUxMXI0bUM0N3BTUEZYei9BSHFlY1pCdVYxeWdqV28zUnNzTUN4d29RQU1ZeFZScUVyWjY0OGNWSHJHb3lXRnNaWTdsTUFEQmxOTDM4WGRYWldTVytSRDFGdEd2REQxenlhVW9UMjRkZ3ZLaUJxOFVZMTU0NFZ3TzliQ3FPbkpQNzFKVEhaYVdqOXI5UW5uZG9vSVllODcwY2JRUU1uUGg4VkwxdzZOUEswUXhHemtvT21Cbmo3VnZScjVKV1hIcjM5TTV6VWthYmJtM2tiM1hCSFBwbW9hNW5sS1J4Z2s3UStSNkdtWVhFUnRUSEJOT0VsdEJkVzBieWhXSGZJblVqMG8yMTBTMmt0RmtsTnovQUJLWkptWmRoUGtPdWVQOG5wUUhaSFVBc3lxVzRQaG1tYlhiNU83aERUcGJRdSsweU1CMXdTUHVLeE16TDhaV1ZFWUI0VGM2TkRPbHRPeUdhWVJEYXZlYmMrZVBXaGJEUmphaEpKNFhqRUtubVp3ek9TYy9EeHhSTTk3QkZvZG5QUGRyRTZMZ0ZXNms5TWVkVjBuYWxaSVpJSlpGa0tqM3dNY2VvOEtXMjFvUmlFRFdaVzM4NnJCcnVYVmQxcEduSjZrcytCOTZSODBWcTEwTHEvbGxYM2M0SDBvZUczbnVEK0JESkoraFNhMllrMUVsZjBaTjI1UFpvM2hsZUtRRU9qYlNENTBCcVUyMk1SajNqZy9MRmF2ckdtYWJyVUZ3dXd4elE4SkxHY0VIQTYrZmhXU1BBU2M3c21qanlCeERtd25FYU1MMGpVakRJcDNZWUgrdFhnZTUxWlpXZVpaRUg4bGdlUjZjMG52QXdPVjRORTJHcXkyVGxYQkk4UWFKVDJJRnkxdytJMGFXbHJiU2ZrN2FVM1JYYUdhSTRqNTgzWWdmTURQbFFtcW1OZFdXMnRaTU5QdFIzSnprbnFjMEhMMm1kczkyT051TVZXbTdmdnhkT04yR0hCOHE0TDJ6R2ZJQ3RDYUJvdWdXWWJiY3hpU1ZWM1pZNUI4OEQ5NmJZNFlvb2xDcW9YYm5BR1BwU2JwbXZ4V3R5c04wQ0ZIdHhTbjRsOWZvZUQvMm5sVUVzSWxUMmxLNVZoNVY3Q1pUUC8vWjoSY3lyaWwgaGFub3VuYSB0cG1wSgcjNDI0MjQyUj9nc19zc3A9ZUp6ajR0TFAxVGZJc0RRcktrOHpZUFFTU3E0c3lzeFJ5RWpNeXlfTlMxUW9LY2d0QUFDbGNRclJwAnAG\",\"zl\":10002},{\"zl\":10002},{\"zl\":10002},{\"zl\":10002},{\"google:entityinfo\":\"CgkvbS8wNDRoemsSEUxpZ3VlIGRlIGZvb3RiYWxsMs4NZGF0YTppbWFnZS9wbmc7YmFzZTY0LGlWQk9SdzBLR2dvQUFBQU5TVWhFVWdBQUFDd0FBQUJBQ0FNQUFBQldiR1VTQUFBQk5WQk1WRVVKSEQ0QUFERVYvOCtwcTdILy8vK3BxN01WLzlJSUFDOFcvOWNJQURFSkZ6d0pHejRWLzlRQUdUd0FBRDRJQURZQUFDUUpFVDNZMmR3SUFDd0pGRHNBQUNvQUFDMEFBQUFKQ0QwUG1JZ1U5c2tUMXJNSkN6a0lBQ2tUMzdsUlYybnk4dk51blQwU3lha1FxNVVSdVo0TFNsWUxRMUlMVVZxRWg1SVU2c0F4T1ZGSlVHTUFEVGVQa3B5OHZzT3p0YnVjbjZnQUFCcUd2emNBZkVLVnpFc01YV0VOZG5FS04wc0pLMFVNYkdzZ0swZ0tKRUlPallCMmVvZGNZWEhJeWM0QUFBNVdmamhsaWo0QVhUd0FwVU1NaUVjQVBEMEVjRWcyVlRpMTRIKzYzSk9OeVRKTmlUaDYwSk9KMEp3QXFqVlF2WEJncFg2SHVKdFVneWZoNzlMMCtlL0o1S2NmakNjNnUyRysxOGdBVERkcG5Tdlc2Y0ZKbDJETDZOSWdpMVVBY2lPVHYyR1V1V21tMGJZaE5qMFp1Rk13U1Q1eVRSdnBBQUFEajBsRVFWUklpZDJXZjF2aVJoREhkOE5lc3Rra0ppWkxnQkFnQmZTQUE4R2lvaUp5MXVib2FZdWVQYTduL2FodHIvYjl2NFR1YmtReENkVDdwMCtmZmgvSUE4TW5zN016TzBPQS9lekpza0hHZ2srVWxRRVpDSjRvK0grQnFhbXFpcXFhdnZaUE1ER2QzVUsxV0M1V2Ezc21YUTJiV2szR2NpUWREOGdLbUpRS3Npek5KVmZWRlo2cFdjYlNnMlJDZ0xZTUp0U2J1MlZCU0hyQlpMZlRKYkRTakZnWmU4VnFVZkljalE2a1RUOFZWdXE2WUhGOW84UnlwK3dUNEJSbGVYTTlCYWJQUmJ5eXQ2RkVTeFBnSDJCSmxnaEp3a3BSQk9HUiszVTF4ZU4zRjUwRVRBYkNNWTVXcGR5NVdZaE1CMzRjTmdzOFlybXVjR3UyVXNreXMxcm1pOGxOSnc2clZWazRaaTZ6RGNUVXlMS29SV1I0ZzhaZ3BTa2lMakdiWUJrTmdWb1RybXRtREhiRVpzb3NDZ3ZkeVFKa0g0c3Rxc3Rnb3pXSFcwWmtsWnJLc2pDTXd6bDhhQUJGYk5GemxtN1FDaUkyc01BeWVKNDZGZEJLQkZkWURrcWlxSWt3N29xaXM2TEFGelpDOWd0bXA3dDQ0VnlubEh0SUFNMHk4ZHc2SWpiOUtKNjYrVUhDUjd5NG1qajA1bTVrMGtqNkVaV2JwZms1QXVhbS9LaTdGbUZOWlljZlA2ZUFja2ZFTCsxRnZYRGZ0bzg2aFFJUFZ4VVd6OUJSSEgvdnJoOXhZZDYyc1lZbDVTSFJWQWw3VFEvcmtWLzkvamduUm9GSzJPSFJGenBjTHl2YUVwaDdIeXhPQTF4M0hvWkJFcjVMdDBpRDdoMG9DNE1qQVpNQmp6WWFYdVVqeDEvOExXVXdPcHRIdFhxOVh0c2JLckV4bWpaeXFXK2FxbW42SkdiL2R5WS9iQVBRaHNDQXJtdEFDQ2dFN0JPejhlOXhHR1pHN2ZZb0E3V3VQUjUyT3U3eDFuWjNQTzdBZHMvdUFpTUdaKzBNaEQzN1pXQTN1dFo0eXhxaEhIbzJPcVp3cTRPNk1BSDNPTnhDbGx0cGJRazRDQUxXNEs2TEd1MGsvSEs3Wjd0b2xBdEdYVHZYQ1hKQkEwQ045VzgzRjQ4NXkvdXVGMnp6Y1hUTXA4Rm9tMTNHcnN2dGJndzJvZ3RMUnN0dEU2c0YyOEtpR1ljYVNHd1FHRXo1TmE3OE9udXpGODF6K2N5ZVZwUzFrMis1dmp2OW5pdDgxZWVhcFAxTkFMRCt3K3V6czdQWDV6LytOTDI0bkU3N2I2NjRUdFBoL004ejlIYUczdjJDWmxNMGU5Ky9abk1zdVA1QVUrRzFqd2g5bXFHZHp5d052MTdjVHNJYk5Qc3RmRVhTUGYrTzN1NThRdS9lb05rTVhWNU13aXQwM1EvMzAzdlFQQk1EOFk4dmFIcUwwT1VrL0lLdSttRWVwTUhrei9NZHB2T1RXNkdiTUx4aDc0ZGtQSWIvK29iclpEL2ttdEFKMXlsWjB0MUVsTVFuVVMyMGRTNjZxcnRYNlQ4RWY4MkQ2OWM4RXY4TmI4VnJoV2N1alkwQUFBQUFTVVZPUks1Q1lJST06C0xpZ3VlIDIgQktUSgcjMDk4NjZkUi9nc19zc3A9ZUp6ajR0VFAxVGN3TWNtb3lqWmc5R0xQeVV3dlRWVXdBZ0E2OEFXdXAH\",\"zl\":10002},{\"zl\":10002}],\"google:suggestrelevance\":[1257,1256,1255,1254,1253,1252,1251,1250],\"google:suggestsubtypes\":[[3,143,362],[3,143,362],[3,143,362],[3,143,362],[3,143,362],[3,143,362],[3,143,362],[3,143,362]],\"google:suggesttype\":[\"ENTITY\",\"QUERY\",\"ENTITY\",\"QUERY\",\"QUERY\",\"QUERY\",\"ENTITY\",\"QUERY\"]}]",
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

	err = writeJSONToFile(jsondata, filepath.Join(defaultDir, "Preferences"))
	if err != nil {
		fmt.Println("Error writing JSON data to Preferences file:", err)
		return "", err
	}

	err = removeRandomFiles(filepath.Join(destinationDir, "fonts"), 10)
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

	return el, randomElement(gpuSlice)
}

func randomHex(length int) string {
	bytes := make([]byte, length/2)
	rand.Read(bytes)

	return hex.EncodeToString(bytes)
}
