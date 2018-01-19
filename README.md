# 위성 Tile들의 해상도를 낮추어 파일 크기를 비교하는 소스입니다.
 MozJpegConverter를 통해 90%, 80%, 70%씩 해상도를 낮추었고 원본과 파일 크기와 사용자 육안으로 봤을 때의 차이점을 비교하였습니다.

## 사용법
### 1. Flag 설명
* -site : 'naver', 'daum', 'vworld'   .. TileMap의 해당 사이트
* -location : 해당 지역 이름

### 2. 실행 방법  
#### 해당 소스 URL
[https://github.com/JinHyukParkk/DownloadTile.git](https://github.com/JinHyukParkk/DownloadTile.git)

#### 2.1. 리눅스, Mac OS X  환경
  1. 'resolutionOfTiles_Comparison' 이라는 실행 파일 다운
  2. 터미널 실행
      * ./resolutionOfTiles_Comparison -site [site] -location [지역명]
  3. xlsx 파일로 저장

#### 2.2. Window 환경
  1. 'resolutionOfTiles_Comparison.exe' 이라는 실행 파일 다운
  2. CMD 관리자 모드로 실행
      * resolutionOfTiles_Comparison.exe -site [site] -location [지역명]
  3. xlsx 파일로 저장

### 3. Golang 라이브러리로 사용
