# 위성 Tile들의 해상도를 낮추어 파일 크기를 비교하는 소스입니다.
 MozJpegConverter를 통해 90%, 80%, 70%씩 해상도를 낮추었고 원본과 파일 크기와 사용자 육안으로 봤을 때의 차이점을 비교하였습니다. 해당 소스는 해상도별 tile들의 size의 max, min, avg 값을 구하는 소스입니다.

## 사용법
### 1. Flag 설명
* -site : 'naver', 'daum', 'vworld'   .. TileMap의 해당 사이트
* -location : 해당 지역 이름
* -type : site가 'vworld'라면 2D or 3D 선택
#### 1.1 타일 데이터 저장 디렉터리 형식
* tileData란 디렉터리 안에 naver_[해상도] 별로 저장 되어있어야 하고, 그 안에 [위치] 디렉터리 안에 타일들이 저장되어있어야 합니다.
```
.tileDate
├── naver_100
|   ├── 독도
|   ├── 북한산
|   └── ...
├── naver_90
├── naver_80
├── naver_70
├── vworld_100
|   ├── 2D독도
|   ├── ...
|   ├── 3D독도
|   └── ...
├── vworld_90
├── vworld_80
└── ...
```

### 2. 실행 방법  
#### 해당 소스 URL
[https://github.com/JinHyukParkk/DownloadTile.git](https://github.com/JinHyukParkk/DownloadTile.git)

#### 2.1. 리눅스, Mac OS X  환경
  1. 'resolutionOfTiles_Comparison' 이라는 실행 파일 다운
  2. 터미널 실행
      * ./resolutionOfTiles_Comparison -site [site] -location [지역명]
  3. result/[site]_[지역]_Result.txt 라는 텍스트 파일에 결과 저장

#### 2.2. Window 환경
  1. 'resolutionOfTiles_Comparison.exe' 이라는 실행 파일 다운
  2. CMD 관리자 모드로 실행
      * resolutionOfTiles_Comparison.exe -site [site] -location [지역명]
  3. result/[site]_[지역]_Result.txt 라는 텍스트 파일에 결과 저장

### 3. 사용자 육안으로 타일 확인
링크 : [https://github.com/IngIeoAndSpare/SortingImageDisplay](https://github.com/IngIeoAndSpare/SortingImageDisplay)
