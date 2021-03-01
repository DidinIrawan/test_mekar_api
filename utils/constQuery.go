package utils

const (
	//Jobs
	GETALLJOB = "select * from tb_pekerjaan where status='A'"
	GETINSERTJOB = "INSERT INTO tb_pekerjaan(pekerjaan) VALUES (?)"
	GETUPDATEJOB = "UPDATE tb_pekerjaan set pekerjaan=? where id_pekerjaan=?"

	//Education
	GETALLEDUCATIONS = "select * from tb_pendidikan where status='A'"
	GETINSERTEDUCATION = "INSERT INTO tb_pendidikan(pendidikan) VALUES (?)"

	//Users
	GETALLUSERS = "select tu.id_user,tu.nama,tu.tanggal_lahir,tu.no_ktp,tp.id_pekerjaan,tp.pekerjaan,tpd.id_pendidikan,tpd.pendidikan,tu.status from tb_user tu join tb_pekerjaan tp on tu.id_pekerjaan=tp.id_pekerjaan join tb_pendidikan tpd on tu.id_pendidikan=tpd.id_pendidikan where tu.status='A';"
	GETUSERBYID ="select tu.id_user,tu.nama,tu.tanggal_lahir,tu.no_ktp,tp.id_pekerjaan,tp.pekerjaan,tpd.id_pendidikan,tpd.pendidikan,tu.status from tb_user tu join tb_pekerjaan tp on tu.id_pekerjaan=tp.id_pekerjaan join tb_pendidikan tpd on tu.id_pendidikan=tpd.id_pendidikan where tu.status='A'& tu.id_user=?"
	GETINSERTUSER = "insert into tb_user(nama,tanggal_lahir,no_ktp,id_pekerjaan,id_pendidikan) values (?,?,?,?,?);"
	GETUPDATEUSERBYID = "update tb_user Set nama=?, tanggal_lahir=?,no_ktp=?,id_pekerjaan=?,id_pendidikan=? where id_user=?"
)