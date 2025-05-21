package main

import (
	"fmt"
	"math"
	"time"
)

func checkLeavePermission(joinDateStr, leaveDateStr string, cutiBersama, durasiCuti int) (bool, string) {
	layout := "2006-01-02"

	joinDate, err := time.Parse(layout, joinDateStr)
	if err != nil {
		return false, "Invalid join date format"
	}

	leaveDate, err := time.Parse(layout, leaveDateStr)
	if err != nil {
		return false, "Invalid leave date format"
	}
	if leaveDate.Before(joinDate) {
		return false, "Leave date cannot be before join date"
	}

	officeLeave := 14
	totalPersonalLeave := officeLeave - cutiBersama

	limitDate := joinDate.AddDate(0, 0, 180)
	if leaveDate.Before(limitDate) {
		return false, "Karena belum 180 hari sejak tanggal join karyawan"
	}

	yearEnd := time.Date(joinDate.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	if leaveDate.Year() > joinDate.Year() {
		yearEnd = time.Date(leaveDate.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	}
	if limitDate.After(yearEnd) {
		return false, "Tidak ada sisa hari kerja untuk cuti di tahun ini"
	}

	validDays := int(yearEnd.Sub(limitDate).Hours()/24) + 1
	eligibleLeave := int(math.Floor(float64(validDays) / 365.0 * float64(totalPersonalLeave)))

	if eligibleLeave <= 0 {
		return false, "Tidak memiliki jatah cuti pribadi"
	}

	if durasiCuti > eligibleLeave {
		return false, fmt.Sprintf("Karena hanya boleh mengambil %d hari cuti", eligibleLeave)
	}

	if durasiCuti > 3 {
		return false, "Cuti pribadi maksimal 3 hari berturut-turut"
	}

	return true, ""
}

func main() {
	testCases := []struct {
		cutiBersama int
		joinDate    string
		leaveDate   string
		durasi      int
	}{
		{7, "2021-05-01", "2021-07-05", 1}, // False, belum 180 hari
		{7, "2021-05-01", "2021-11-05", 3}, // False, hanya dapat 1 hari
		{7, "2021-01-05", "2021-12-18", 1}, // True
		{7, "2021-01-05", "2021-12-18", 3}, // True
	}

	for _, t := range testCases {
		allowed, reason := checkLeavePermission(t.joinDate, t.leaveDate, t.cutiBersama, t.durasi)
		fmt.Printf("Input: \n - Jumlah Cuti Bersama = %d", t.cutiBersama)
		fmt.Printf("\n - Tanggal join karyawan = %s", t.joinDate)
		fmt.Printf("\n - Tanggal rencana cuti = %s", t.leaveDate)
		fmt.Printf("\n - Durasi cuti (hari) = %d", t.durasi)
		fmt.Printf("\nOutput: \n %v", allowed)

		if !allowed {
			fmt.Printf("\n Alasan: %s", reason)
		}
		fmt.Println()
	}
}
