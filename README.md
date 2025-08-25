# Password Generator

Go 언어로 작성된 안전한 비밀번호 해시 생성기입니다. 이 프로그램은 사용자 입력을 받아 암호학적으로 안전한 salt와 해시를 생성합니다.

## 주요 기능

- **랜덤 Salt 생성**: 16바이트(32자리 16진수) 랜덤 salt 생성
- **SHA-256 해싱**: SHA-256 해시 알고리즘 사용
- **Key Stretching**: 10,000회 반복 해시 계산으로 보안 강화
- **JSON 저장**: 생성된 데이터를 JSON 형식으로 파일 저장

## 보안 특징

- `crypto/rand` 패키지를 사용한 암호학적으로 안전한 랜덤 생성
- Salt를 통한 Rainbow Table 공격 방지
- Key Stretching을 통한 Brute Force 공격 저항성 향상
- SHA-256 해시 알고리즘 사용

## 설치 및 실행

### 요구사항
- Go 1.16 이상

### 실행 방법
```bash
# 프로젝트 디렉토리로 이동
cd password_generator

# 프로그램 실행
go run main.go
```

### 빌드 방법
```bash
# 실행 파일 빌드
go build -o password_generator main.go

# 실행
./password_generator
```

## 사용법

1. 프로그램을 실행합니다
2. "Enter the original data:" 프롬프트가 나타나면 원본 데이터를 입력합니다
3. 프로그램이 자동으로 salt를 생성하고 해시를 계산합니다
4. 결과가 콘솔에 출력되고 `password_data.json` 파일에 저장됩니다

## 출력 예시

```
Enter the original data: mypassword123
Salt: a5456ad97099ccbe68a4063df914513e
Hashed value: 36d43b764762822a4b4b104cf8da22656b468da5464d87061b802323a5b24032
Success to save data
```

## 생성되는 파일

### password_data.json
```json
{
  "plain_text": "원본 데이터",
  "salt": "생성된 salt",
  "hashed": "해시된 값"
}
```

## 코드 구조

- `generateRandomSalt(length int)`: 랜덤 salt 생성
- `hashPlainTextWithSaltAndStretching(plainText, salt, iterations)`: salt와 key stretching을 사용한 해시 생성
- `saveToFile(filename string, data PasswordData)`: JSON 파일로 데이터 저장
- `main()`: 메인 프로그램 로직

## 보안 고려사항

- 이 프로그램은 교육 및 테스트 목적으로 설계되었습니다
- 프로덕션 환경에서는 추가적인 보안 검토가 필요합니다
- 생성된 해시는 단방향이므로 원본 데이터로 복원할 수 없습니다

## 라이선스

이 프로젝트는 MIT 라이선스 하에 배포됩니다.

## 기여

버그 리포트나 기능 제안은 이슈를 통해 제출해 주세요.
