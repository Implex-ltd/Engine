package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var (
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	GpuData    = map[string][]string{
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

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomBrowserSize(minSize, maxSize int) (width, height int64) {
	aspectRatios := []float64{16.0 / 9.0, 16.0 / 10.0, 4.0 / 3.0}

	rand.Shuffle(len(aspectRatios), func(i, j int) {
		aspectRatios[i], aspectRatios[j] = aspectRatios[j], aspectRatios[i]
	})

	aspectRatio := aspectRatios[0]

	width = int64(rand.Intn(maxSize-minSize+1) + minSize)
	height = int64(float64(width) / aspectRatio)

	return width, height
}

func RandomElement(slice []string) string {
	index := rand.Intn(len(slice))
	return slice[index]
}

func RandomKey(m map[string][]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys[rand.Intn(len(keys))]
}

func GetRandomVendorRenderer() (string, string) {
	el := RandomKey(GpuData)
	gpuSlice := GpuData[el]

	if len(gpuSlice) == 0 {
		return "", ""
	}

	return el, RandomElement(gpuSlice)
}

func RandomElementInt(slice []int64) int64 {
	index := rand.Intn(len(slice))
	return slice[index]
}

func GenerateRandomMD5() string {
	h := md5.New()
	_, _ = io.WriteString(h, RandStringBytes(32))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CopyDir(src, dest string) error {
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
			if err := CopyDir(srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFile(src, dest string) error {
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

func WriteJSONToFile(data []byte, filePath string) error {
	return os.WriteFile(filePath, data, 0644)
}

func RemoveRandomFiles(dirPath string, numFiles int) error {
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

func GenerateRandomIP() string {
	octet1 := rand.Intn(256)
	octet2 := rand.Intn(256)
	octet3 := rand.Intn(256)
	octet4 := rand.Intn(256)

	ipAddress := fmt.Sprintf("%d.%d.%d.%d", octet1, octet2, octet3, octet4)

	return ipAddress
}

func RandomHex(length int) string {
	bytes := make([]byte, length/2)
	rand.Read(bytes)

	return hex.EncodeToString(bytes)
}
