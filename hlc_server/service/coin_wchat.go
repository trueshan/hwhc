package service

import (
	"fmt"
	"github.com/hwhc/hlc_server/hoo"
	"github.com/hwhc/hlc_server/log"
	"github.com/hwhc/hlc_server/mysql"
	"github.com/hwhc/hlc_server/persistence"
	"strconv"
)

func Transfer_app(userId int64, transferId int64, cid int64, adminid string) (int, string) {

	//审核开关
	auditCoinSwitch := persistence.GetConfig(mysql.Get(),persistence.AuditCoinKey)
	if auditCoinSwitch != "on"{
		return -2012, "系统维护中....."
	}

	xmysql := mysql.Begin()
	defer xmysql.Commit()
	// 验证余额    验证余额    验证余额
	trasfer := persistence.TransferbyId(mysql.Get(), userId, transferId, cid)
	log.Info(fmt.Sprintf("[debug] Transfer_app trasfer->trasferId：%d,transferId:%d,cid:%d,adminid : %v", trasfer.Id, transferId, cid,adminid))
	if !persistence.UpdateTransferStatus(xmysql, trasfer.Tx_hash, trasfer.UserId) {
		return -2011, "更新订单状态失败，刷新后重试"
	}

	coinInfo := GetcoinbyId(cid)
	amount := strconv.FormatFloat(0-trasfer.Amount, 'f', 6, 64)
	hooOrderNo, msg := hoo.TransferHCOut(trasfer.Tx_hash, amount, coinInfo.Coinname, trasfer.Address, coinInfo.ContractCddress, coinInfo.Tokenname, trasfer.Memo)

	persistence.UpdateTransferHooOrderNo(xmysql, trasfer.Id, trasfer.UserId, hooOrderNo)
	log.Info(fmt.Sprintf("[debug] 管理员：%v,处理了订单：%d ,hooOrderNo：%s", adminid, transferId, hooOrderNo))

	if hooOrderNo == "" {
		xmysql.Rollback()
		log.Info(fmt.Sprintf("[debug]transfer rollback,hooOrderNo:%s",hooOrderNo))
		return -10010, msg
	} else {
		log.Info("[debug]transfer success")
		return 0, ""
	}
}


func Transfer_check(userId int64, transferId int64, cid int64, adminid string) (int, string) {

	trasfer := persistence.TransferbyId(mysql.Get(), userId, transferId, cid)
	if trasfer.Id == 0{
		return -10012,"该笔交易不存在，刷新后重试"
	}
	if trasfer.Tx_status == -1{
		log.Info(fmt.Sprintf("[debug]Transfer_check status is -1 transferId:%d ,cid :%d,adminid:%s",transferId,cid,adminid))
		return -10012,"该笔交易已被驳回，刷新后重试"
	}
	if trasfer.Tx_status == 1{
		log.Info(fmt.Sprintf("[debug]Transfer_check status is 1 transferId:%d ,cid :%d,adminid:%s",transferId,cid,adminid))
		return 0,""
	}

	//查询hoo钱包状态
	fooOrder := hoo.GetOrder(trasfer.Tx_hash)
	if fooOrder.Data.OuterOrderNo == "" || fooOrder.Data.OrderNo == "" {
		return -10012,"该笔交易在hoo未查询到"
	}

	hooOrderNo := fooOrder.Data.OrderNo

	if fooOrder.Data.Status == "success"{
		xmysql := mysql.Begin()

		//更新状态
		if !persistence.UpdateTransferStatus(xmysql, trasfer.Tx_hash, trasfer.UserId) {
			xmysql.Rollback()
			log.Info(fmt.Sprintf("[debug] 管理员：%s,Transfer_check 更新订单状态失败 transferId ：%d ,hooOrderNo：%s ,userid:%d , cid:%d ", adminid, transferId, hooOrderNo,userId,cid))
			return -2011, "更新订单状态失败，刷新后重试"
		}
		//更新hoo地址
		if !persistence.UpdateTransferHooOrderNo(xmysql, trasfer.Id, trasfer.UserId, hooOrderNo){
			xmysql.Rollback()
			log.Info(fmt.Sprintf("[debug] 管理员：%s,Transfer_check 更新hooOrderNo失败 transferId ：%d ,hooOrderNo：%s ,userid:%d , cid:%d ", adminid, transferId, hooOrderNo,userId,cid))
			return -10010, "更新hooOrderNo失败，刷新后重试"
		}

		log.Info("[debug]Transfer_check success")
		xmysql.Commit()
		return 0, ""
	}else{
		return -10012,"该笔交易在hoo交易未完成，刷新后重试"
	}

}
