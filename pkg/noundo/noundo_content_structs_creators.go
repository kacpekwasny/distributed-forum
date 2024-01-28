package noundo

import "time"

func CreateUserInfo(user UserIdentityIface, usingHistoryName string) UserInfo {
	return UserInfo{
		username:       user.GetUsername(),
		FUsername:      user.GetFUsername(),
		parentServer:   user.GetParentServerName(),
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
