package machinary

type Machinary struct {
	MachinaryID         int64   `json:"machineID"`
	CategoryID          int64   `json:"categoryID"`
	OwnerID             int64   `json:"ownerID"`
	MachinaryName       string  `json:"machinename"`
	MachinaryImage      string  `json:"image"`
	MachinaryImageLarge string  `json:"imagelarge"`
	Description         string  `json:"description"`
	PackingType         string  `json:"packingtype"`
	Stock               int     `json:"stock"`
	Price               float64 `json:"price"`
	PriceString         string  `json:"priceString"`
}
type MachinarysList struct {
	Machinarys []Machinary `json:"machines"`
	StatusCode int         `json:"status"`
}

type Catergory struct {
	CatergoryID    int64  `json:"categoryid"`
	CatergoryName  string `json:"categoryname"`
	CatergoryImage string `json:"image"`
}
type CategoryLists struct {
	Categories []Catergory `json:"categories"`
	StatusCode int         `json:"status"`
}

// func (p *Machinary) AddMachinary() error {
// 	res, err := database.DB.Exec(`
// 	INSERT INTO machines
// 	(category_id,owner_id,machine_name,machine_image,machine_image_large,descriptions,price,stock,machine_packtype)
// 	VALUES (?,?,?,?,?,?,?,?,?)`, p.CategoryID, p.OwnerID, p.MachinaryName, p.MachinaryImage,
// 		p.MachinaryImageLarge, p.Description, p.Price, p.Stock, p.PackingType)
// 	if err != nil {
// 		return err
// 	}
// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		return errors.New("could not get the latest id")
// 	}
// 	p.MachinaryID = id
// 	return nil
// }
// func (p *Machinary) UpdateMachinary() error {
// 	res, err := database.DB.Exec(`UPDATE machines SET
// 	category_id=?,
// 	owner_id=?,
//     machine_name=?,
//     machine_image=?,
//     machine_image_large=?,
//     descriptions=?,
//     price=?,
//     stock=?,
//     machine_packtype=?
// 	WHERE machineID=?;`, p.CategoryID, p.OwnerID, p.MachinaryName, p.MachinaryImage,
// 		p.MachinaryImageLarge, p.Description, p.Price, p.Stock, p.PackingType, p.MachinaryID)
// 	if err != nil {
// 		return err
// 	}
// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if count == 0 {
// 		return errors.New("the machine you wish to update does not exist")
// 	}
// 	return nil
// }
// func (p *Machinary) DeleteMachinary() error {
// 	_, err := database.DB.Exec("DELETE FROM machines WHERE machineID=?", p.MachinaryID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (p *Machinary) GetCurrentStock() int {
// 	tempStock := 0
// 	err := database.DB.QueryRow("SELECT stock FROM machines WHERE machineID = ?", p.MachinaryID).Scan(&tempStock)
// 	if err != nil {
// 		fmt.Printf("Unable to retrive current stock because of error %v", err)
// 	}
// 	p.Stock = tempStock
// 	return p.Stock
// }

// //Database transactions mehtod that is dependat if the invoice was placed succefully
// func (p *Machinary) DeceremtStockBy(tx *sql.Tx, reductionAmount int64) error {
// 	currentStock := p.GetCurrentStock()
// 	newStock := currentStock - int(reductionAmount)
// 	if !p.CanBePurchased(reductionAmount) {
// 		return fmt.Errorf("the stock (%v) being purchased is relatively higher than the available stock(%v)", currentStock, reductionAmount)
// 	}
// 	_, err := tx.Exec("UPDATE machines SET stock=? WHERE machineID=?", newStock, p.MachinaryID)
// 	if err != nil {
// 		return fmt.Errorf("---%v", err)
// 	}
// 	p.Stock = newStock
// 	return nil
// }

// func (p *Machinary) CanBePurchased(quantity int64) bool {
// 	currentStock := p.GetCurrentStock()
// 	newStock := currentStock - int(quantity)
// 	return newStock > 0
// }

// func (p *Machinary) GetWalletOfMachinaryOwner() *wallet.Wallet {
// 	ownersPhonenumber := ""
// 	err := database.DB.QueryRow("SELECT phonenumber FROM machines INNER JOIN user_registration ON machines.owner_id = user_registration.userid WHERE machineID=? ", p.MachinaryID).Scan(&ownersPhonenumber)
// 	if err != nil {
// 		return nil
// 	}
// 	return wallet.GetWalletByAddress(ownersPhonenumber)
// }

// func (p *Machinary) GetMachinaryShortName() string {
// 	//IMplement Name shortenning here to be able to not fill a users transaction with long unecessary names
// 	return p.MachinaryName
// }

// func GetMachinarysByOwnerID(owner_id int64) (*MachinarysList, error) {
// 	result := new(MachinarysList)
// 	rows, err := database.DB.Query(`
// 	SELECT machineID,
//   	category_id,
// 	machine_name,
// 	machine_image,
// 	machine_image_large,
// 	descriptions,
// 	price,stock,
// 	machine_packtype
// 	FROM machines WHERE owner_id=?;
// 	`, owner_id)

// 	if err != nil {
// 		result.StatusCode = -500
// 		return result, err
// 	}

// 	for rows.Next() {
// 		singleMachinary := Machinary{}
// 		if err := rows.Scan(
// 			&singleMachinary.MachinaryID,
// 			&singleMachinary.CategoryID,
// 			&singleMachinary.MachinaryName,
// 			&singleMachinary.MachinaryImage,
// 			&singleMachinary.MachinaryImageLarge,
// 			&singleMachinary.Description,
// 			&singleMachinary.Price,
// 			&singleMachinary.Stock,
// 			&singleMachinary.PackingType); err != nil {
// 			result.StatusCode = -501
// 			return result, err
// 		}
// 		result.Machinarys = append(result.Machinarys, singleMachinary)
// 	}
// 	if err != nil {
// 		result.StatusCode = -502
// 		return result, err
// 	}
// 	if result.Machinarys == nil {
// 		result.StatusCode = -503
// 		result.Machinarys = []Machinary{}
// 	}
// 	defer rows.Close()
// 	return result, nil
// }

// func GetMachinarysByCategoryID(category_id int64) (*MachinarysList, error) {
// 	result := new(MachinarysList)
// 	rows, err := database.DB.Query(`
// 	SELECT machineID,
// 	owner_id,
//   	category_id,
// 	machine_name,
// 	machine_image,
// 	machine_image_large,
// 	descriptions,
// 	price,stock,
// 	machine_packtype
// 	FROM machines WHERE category_id=?;
// 	`, category_id)

// 	if err != nil {
// 		result.StatusCode = -500
// 		return result, err
// 	}

// 	for rows.Next() {
// 		singleMachinary := Machinary{}
// 		if err := rows.Scan(
// 			&singleMachinary.MachinaryID,
// 			&singleMachinary.OwnerID,
// 			&singleMachinary.CategoryID,
// 			&singleMachinary.MachinaryName,
// 			&singleMachinary.MachinaryImage,
// 			&singleMachinary.MachinaryImageLarge,
// 			&singleMachinary.Description,
// 			&singleMachinary.Price,
// 			&singleMachinary.Stock,
// 			&singleMachinary.PackingType); err != nil {
// 			result.StatusCode = -501
// 			return result, err
// 		}
// 		result.Machinarys = append(result.Machinarys, singleMachinary)
// 	}
// 	if err != nil {
// 		result.StatusCode = -502
// 		return result, err
// 	}
// 	if result.Machinarys == nil {
// 		result.StatusCode = -503
// 		result.Machinarys = []Machinary{}
// 	}
// 	defer rows.Close()
// 	return result, nil
// }

// func GetCategories() (*CategoryLists, error) {
// 	result := new(CategoryLists)
// 	rows, err := database.DB.Query("SELECT category_id,category_name,category_image FROM m_categories")
// 	if err != nil {
// 		result.StatusCode = -500
// 		return result, err
// 	}

// 	for rows.Next() {
// 		singleCategory := Catergory{}
// 		if err := rows.Scan(&singleCategory.CatergoryID, &singleCategory.CatergoryName, &singleCategory.CatergoryImage); err != nil {
// 			result.StatusCode = -501
// 			return result, err
// 		}
// 		result.Categories = append(result.Categories, singleCategory)
// 	}
// 	if err != nil {
// 		result.StatusCode = -502
// 		return result, err
// 	}
// 	//To avoid passing null back to the user
// 	if result.Categories == nil {
// 		result.StatusCode = -503
// 		result.Categories = []Catergory{}
// 	}
// 	defer rows.Close()
// 	return result, nil
// }
// func GetMachinaryByMachinaryID(machineID int64) *Machinary {
// 	tempMachinary := new(Machinary)
// 	database.DB.QueryRow(`
// 	SELECT machineID,
// 	owner_id,
//   	category_id,
// 	machine_name,
// 	machine_image,
// 	machine_image_large,
// 	descriptions,
// 	price,stock,
// 	machine_packtype
// 	FROM machines WHERE machineID=?;
// 	`, machineID).Scan(&tempMachinary.MachinaryID,
// 		&tempMachinary.OwnerID,
// 		&tempMachinary.CategoryID,
// 		&tempMachinary.MachinaryName,
// 		&tempMachinary.MachinaryImage,
// 		&tempMachinary.MachinaryImageLarge,
// 		&tempMachinary.Description,
// 		&tempMachinary.Price,
// 		&tempMachinary.Stock,
// 		&tempMachinary.PackingType)

// 	return tempMachinary
// }
