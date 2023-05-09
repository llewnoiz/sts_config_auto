import { spawn } from 'child_process';


// 1. aws cli 프로파일 생성 
// aws configure --profile jenga


/**
 * Params
 * 
 * mfa arn 정보  : aws 에서mfa 설정후 mfa arn 정보.
 * mfa code 정보 : aws에서 mfa설정 후 mobile 에서 otp code 확인 가능하다.
 * 
 */

// 2. aws mfa 설정
// aws sts get-session-token --serial-number arn:aws:iam::144149479695:mfa/bespin-ci-song --token-code 425387

