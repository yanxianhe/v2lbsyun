.PHONY: install build clean

install:
	pip install nuitka -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com
	pip install -r requirements.txt -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com

build: install
	python3 -m nuitka --onefile --output-filename=v2lbsyun --onefile-tempdir-spec=/tmp/v2lbsyun_{TIME} main.py

clean:
	rm -rf nuitka-crash-report.xml
	rm -f v2lbsyun
	rm -rf /tmp/v2lbsyun_*
