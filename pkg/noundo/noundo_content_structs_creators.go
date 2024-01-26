package noundo

import "time"

func CreateUserInfo(user UserPublicIface, usingHistoryName string) UserInfo {
	return UserInfo{
		Username:       user.Username(),
		FUsername:      user.FullUsername(),
		ParentServer:   user.ParentServerName(),
		UserProfileURL: ProfileURL(user, usingHistoryName),
	}
}

func CreateTimeStamp() TimeStampable {
	return TimeStampable{
		Timestamp: UnixTimeNow(),
	}
}

func UnixTimeNow() int64 {
	return time.Now().Unix()
}
