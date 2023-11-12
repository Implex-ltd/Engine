import httpx, threading, time, itertools

__jwt__ = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJmIjowLCJzIjoyLCJ0IjoidyIsImQiOiJTOFlmVVYwcVk0MUl1dFhLb3dCS1NqV25CZDNpQmwybE1pUDh1THE1L21pdjZYdDRoSnBLREpla3c3dm1odC9wRlFIQ2ZIbFhHSFdLWjZ6aGw4Y0tCRW5MU3lMbUx3emZMUFZSajA3cWk3MlNnWitldjdXd3pIWW9pTHZ3MVFXN3lwZmtmV0xEb1pCNU9jczd3eXlQclhrTTIwTlVHSENXVzV3cTVnSG5ZcWtDNWhNanIrTU1KZlNUUUE9PWlPNnBNOXlaSkFveW4rRVAiLCJsIjoiaHR0cHM6Ly9uZXdhc3NldHMuaGNhcHRjaGEuY29tL2MvYTkxMjcyYSIsImUiOjE2OTI3MjE0NjksIm4iOiJoc3ciLCJjIjoxMDAwfQ.Sf7U9mcIevLIhyza25NLjKuX0VnxfEKve77jFaUjpIBnhPqRTe4gEtkJKptX5BcXx2qMhD0vnwxcymHSmt3j_7pZENnrl1qaU_zinsjt5zypW4HfN2csb_GjjO9vNjcoUdXvTuNakbWl8-kPGlIcgcYohWWYakT7BT6mJKm-7VE"
__iter__ = 100000

response_times = []

urls = itertools.cycle([
      "http://127.0.0.1:1234/n",
      "http://127.0.0.1:1235/n",
      "http://127.0.0.1:1236/n",
      "http://127.0.0.1:1237/n",
      "http://127.0.0.1:1238/n",
      "http://127.0.0.1:1239/n",
      "http://127.0.0.1:1240/n",
])

def test():
    global response_times

    while len(response_times) < __iter__:
        start = time.time()
        n = httpx.post(
            next(urls),
            timeout=5000,
            json={"jwt": __jwt__},
        ).text
        """ n = httpx.get(
            f"http://127.0.0.1:4000/n?req={__jwt__}",
            timeout=None,
        ).text"""
        end = time.time() - start

        response_times.append(end)

        avg_response_time = sum(response_times) / len(response_times)
        print(f"(avg: {avg_response_time:.6f}) took: {end} | {n[:50]}")


if __name__ == "__main__":
	t = time.time()
	for _ in range(100):
		threading.Thread(target=test).start()
	print(time.time() - t)