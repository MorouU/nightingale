SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

CREATE DATABASE IF NOT EXISTS `ni_v1_0` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `ni_v1_0`;

DROP TABLE IF EXISTS `ni_admin_login`;
CREATE TABLE `ni_admin_login` (
  `adminID` int(11) NOT NULL,
  `adminName` varchar(128) NOT NULL,
  `adminPasswd` varchar(64) NOT NULL,
  `adminSalt` varchar(48) NOT NULL,
  `adminEmail` varchar(128) NOT NULL,
  `adminPhone` varchar(20) NOT NULL,
  `lastLoginIP` varchar(32) NOT NULL,
  `lastLoginTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `ni_admin_login` (`adminID`, `adminName`, `adminPasswd`, `adminSalt`, `adminEmail`, `adminPhone`, `lastLoginIP`, `lastLoginTime`) VALUES
(1, 'admin', 'a1cb6b4e78a5132aa340d3e902e9073f959d2b2343e9827e89d104a17baca346', '2b9fc0ff-9439-dad0-b9ad-a895d63a6e46', 'admin@admin.com', '13145678912', '127.0.0.1', '2021-12-31 13:47:39');

DROP TABLE IF EXISTS `ni_app_public_modules`;
CREATE TABLE `ni_app_public_modules` (
  `moduleID` bigint(20) NOT NULL,
  `moduleName` varchar(64) NOT NULL,
  `moduleData` text NOT NULL,
  `createTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `ni_app_public_modules` (`moduleID`, `moduleName`, `moduleData`, `createTime`) VALUES
(0, '默认XSS模块', '支持功能：\r\n* 获取对方当前URL \r\n* 获取对方来路URL \r\n* 获取对方Cookie数据 \r\n* 获取对方操作系统 \r\n* 获取对方浏览器信息 \r\n* 获取对方屏幕分辨率 \r\n* 获取对方网页内容 (会以json格式发送至邮箱) \r\n* 获取对方网页截图 (会以json格式发送至邮箱) ', '2021-04-16 17:57:32'),
(1, '域内请求XSS模块', '支持功能：\r\n* 向指定域内URL进行GET、POST请求和上传文件请求', '2021-05-13 14:01:42'),
(2, 'flash钓鱼XSS模块', '支持功能：\r\n* 显示需要安装flash插件，诱导用户点击并下载文件', '2021-05-14 14:33:38'),
(3, '端口扫描XSS模块', '支持功能：\r\n* 可以探测指定IP的指定范围端口', '2021-05-17 16:03:05');

DROP TABLE IF EXISTS `ni_app_tasks`;
CREATE TABLE `ni_app_tasks` (
  `taskID` bigint(20) NOT NULL,
  `taskModuleID` bigint(20) NOT NULL,
  `taskUserID` int(11) NOT NULL,
  `taskName` varchar(64) NOT NULL,
  `taskModuleName` varchar(64) NOT NULL,
  `taskModulePublic` int(2) NOT NULL,
  `taskData` text NOT NULL,
  `taskCode` text NOT NULL,
  `taskApi` varchar(64) NOT NULL,
  `taskParams` text NOT NULL,
  `taskStatus` int(2) NOT NULL,
  `taskRecordNum` bigint(20) NOT NULL,
  `filePath` text,
  `taskCreateTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `ni_app_task_records`;
CREATE TABLE `ni_app_task_records` (
  `recordID` bigint(20) NOT NULL,
  `taskID` bigint(20) NOT NULL,
  `moduleID` bigint(20) NOT NULL,
  `userID` int(11) NOT NULL,
  `modulePublic` int(2) NOT NULL,
  `getMethod` varchar(16) NOT NULL,
  `getIP` varchar(32) NOT NULL,
  `getResult` text NOT NULL,
  `getModuleResult` text NOT NULL,
  `getTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `ni_user_custom_modules`;
CREATE TABLE `ni_user_custom_modules` (
  `moduleID` bigint(20) NOT NULL,
  `moduleName` varchar(64) NOT NULL,
  `moduleData` text NOT NULL,
  `moduleCode` text NOT NULL,
  `belongUserID` int(11) NOT NULL,
  `createTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `ni_user_forget_auth_code`;
CREATE TABLE `ni_user_forget_auth_code` (
  `email` varchar(128) NOT NULL,
  `activeCode` varchar(64) NOT NULL,
  `time` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `ni_user_login`;
CREATE TABLE `ni_user_login` (
  `userID` int(11) NOT NULL,
  `userName` varchar(128) DEFAULT NULL,
  `userPasswd` varchar(64) DEFAULT NULL,
  `userSalt` varchar(48) DEFAULT NULL,
  `userEmail` varchar(128) DEFAULT NULL,
  `userPhone` varchar(20) DEFAULT NULL,
  `userLevel` int(11) NOT NULL,
  `userCreateTime` datetime DEFAULT NULL,
  `userLoginIP` varchar(32) DEFAULT NULL,
  `lastLoginTime` datetime DEFAULT NULL,
  `active` int(2) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


ALTER TABLE `ni_admin_login`
  ADD PRIMARY KEY (`adminID`);

ALTER TABLE `ni_app_public_modules`
  ADD PRIMARY KEY (`moduleID`);

ALTER TABLE `ni_app_tasks`
  ADD PRIMARY KEY (`taskID`);

ALTER TABLE `ni_app_task_records`
  ADD PRIMARY KEY (`recordID`);

ALTER TABLE `ni_user_custom_modules`
  ADD PRIMARY KEY (`moduleID`);

ALTER TABLE `ni_user_login`
  ADD PRIMARY KEY (`userID`);


ALTER TABLE `ni_admin_login`
  MODIFY `adminID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

ALTER TABLE `ni_app_public_modules`
  MODIFY `moduleID` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

ALTER TABLE `ni_app_tasks`
  MODIFY `taskID` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;

ALTER TABLE `ni_app_task_records`
  MODIFY `recordID` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;

ALTER TABLE `ni_user_custom_modules`
  MODIFY `moduleID` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;

ALTER TABLE `ni_user_login`
  MODIFY `userID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
