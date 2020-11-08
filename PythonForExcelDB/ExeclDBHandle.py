# -*- coding:utf-8 -*-
import os
import pymssql
import xlrd
import xlwt
import datetime
import time
import chardet
import shutil

#数据库全局连接对象
conn=pymssql.connect(host="127.0.0.1",user="sa",password="4869Ahui...A",database="GoldControlDB")
#conn=pymssql.connect(host="47.91.154.143:1433",user="sa",password="4869Ahui...A",database="GoldControlDB")
cou=conn.cursor()

ExcelFilesNumber = 0
ExcelFilesDstFolder = r"E:\\PythonDataProcess\\MassGameDataHandle\\Excel文件\\"

#TargetFolder = r"E:\\PythonDataProcess\\MassGameDataHandle\\888444\\073数据"
#TargetFolder = r"E:\\PythonDataProcess\\MassGameDataHandle\\888444\\769棋牌数据"
TargetFolder = r"E:\\PythonDataProcess\\MassGameDataHandle\\888444"
ExeclDestFile = r"E:\\PythonDataProcess\\MassGameDataHandle\\去重后的数据\\去重后数据_{}.xlsx"
TxtDestFile = r"E:\\PythonDataProcess\\MassGameDataHandle\\去重后的数据\\去重后数据Txt\\{}.txt"

#写数据库
def Write2DB(mobile):
	#插入数据
	sql="INSERT INTO tb_UserMobile(Mobile)VALUES({})".format(mobile)
	cou.execute(sql)
	return

def WriteExcel():
	couts = 0
	sheets_excel = 0
	sheets_text = 0
	books = 0
	workbook = xlwt.Workbook()  			 #新建一个工作簿
	sheet = workbook.add_sheet(str(sheets_excel))  #在工作簿中新建一个表格
	txt_file = open(TxtDestFile.format(sheets_text), 'w')
	
	for mobile in cou.fetchall():
		#print(mobile[0])
		mobile_str = str(mobile[0])
		sheet.write(couts, 0, mobile_str.lstrip('86'))#像表格中写入数据(对应的行和列)
		txt_file.write(mobile_str.lstrip('86'))
		txt_file.write('\n')
		couts += 1
		if couts % 50000 == 0:
			txt_file.close()
			workbook.save(ExeclDestFile.format(books))#保存工作簿
			print("数据写入成功:", sheets_text)
			sheets_text += 1
			sheets_excel += 1
			couts = 0
			if sheets_excel % 20 == 0:
				sheets_excel = 0
				books += 1
				workbook = xlwt.Workbook()  			 #新建一个工作簿
			sheet = workbook.add_sheet(str(sheets_excel))  #在工作簿中新建一个表格
			txt_file = open(TxtDestFile.format(sheets_text), 'w')
	return 
	
#读取数据
def ReadTable():
	sql="SELECT Mobile FROM AWeiTest"
	cou.execute(sql)

	return
	
#提交数据库
def CommitDB():
	conn.commit()
	return
	
#关闭数据库
def CloseDB():
	conn.close()
	return

def ExeclHandle():
	excel_file=xlrd.open_workbook(r'UserIDs.xlsx')
	#获取目标EXCEL文件sheet名
	sheet=excel_file.sheet_by_name('UserID1')
	#打印sheet的名称，行数，列数
	print('表', sheet.name, '共', sheet.nrows,'行', sheet.ncols, '列')
	row_number = sheet.nrows
	index = 1
	while index < row_number:
		row_content=sheet.row_values(index)
		#print("写入号码:", int(row_content[0]))
		Write2DB(int(row_content[0]))
		index += 1
	CommitDB()
	CloseDB()
	
	return 

def GetEncodeMode(file_name):
	bytes = min(32, os.path.getsize(file_name))
	raw_file = open(file_name, 'rb').read(bytes)
	result = chardet.detect(raw_file)
	encoding = result['encoding']
	return encoding

def HandleTxtFile(file_name):
	print("Txt处理文件:", file_name)
	file_encode = GetEncodeMode(file_name)
	try:
		file_handle = open(file_name, encoding = file_encode)
		lines = file_handle.readlines()
		for line in lines:
			#在这里写DB,每个文件提交一次
			#Write2DB(line.strip())
			break
	except UnicodeDecodeError:
		print("		文件处理编码错误:", file_name)
	except Exception as e:
		print("		文件处理未知错误:", file_name)
		print("		", e)
	finally:
		CommitDB()
		file_handle.close()

	return 

def HandleExcelFile(file_name):
	#拷贝到一个目录下统一处理
	print("Execl处理文件:", file_name)
	global ExcelFilesNumber
	ExcelFilesNumber += 1
	target_name = os.path.basename(file_name)
	target_name = str(ExcelFilesNumber) + target_name
	try:
		shutil.copyfile(file_name, ExcelFilesDstFolder + target_name)
	except shutil.SameFileError:
		print("		同名文件异常:", file_name)
		
	return

#判断文件类型并分发给相应文件处理函数
def DispatchFile(file_name):
	extention = os.path.splitext(file_name)[-1][1:]
	if 'txt' ==  extention:
		HandleTxtFile(file_name)
	elif 'xlsx' ==  extention:
		HandleExcelFile(file_name)
	else:
		print("		", extention, "文件不处理", file_name)
	return

def ListDirs():
	print("List Directory")
	Infos = os.walk(r"E:\\PythonDataProcess\\MassGameDataHandle\\QX资料")
	for root,dirs,files in Infos:
		for dir in dirs:
			#获取目录的名称
			print(dir)
			#获取目录的路径
			print(os.path.join(root,dir))

def ClearDstFolder():
	shutil.rmtree(ExcelFilesDstFolder)
	os.mkdir(ExcelFilesDstFolder)
	return 

def PumpIntoDB():
	ClearDstFolder()
	print("List Files")
	Infos = os.walk(TargetFolder)
	for root,dirs,files in Infos:
		for file in files:
			#获取文件所属目录
			#print(root)
			#获取文件路径
			DispatchFile(os.path.join(root,file))
			#break
		#break
	return 
	
def PumpIntoFile():
	ReadTable()
	WriteExcel()
	return 
	
def FilesProcess():
	ClearDstFolder()
	PumpIntoDB()
	PumpIntoFile()
		
if __name__ == "__main__":
	print("Test Begin\n")
	sheet_name = str(1)
	starttime = datetime.datetime.now()
	#FilesProcess()
	PumpIntoFile()
	endtime = datetime.datetime.now()
	print("耗时(s):", (endtime - starttime).seconds)
	
	print("\nTest Done")
	
	
	
	
	

	
