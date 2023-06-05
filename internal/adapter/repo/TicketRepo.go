package repo

import (
	"errors"
	"ticket/goapi/internal/core/port"

	"gorm.io/gorm"
)

type ticketRepo struct {
	db *gorm.DB
}

func NewTicketRepo(db *gorm.DB) port.TicketRepo {
	return ticketRepo{
		db: db,
	}
}

func (c ticketRepo) GetAll() ([]port.Ticket, error) {
	tickets := []port.Ticket{}
	err := c.db.Find(&tickets).Error
	// err := c.db.Order("ticket_id desc").Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}
func (c ticketRepo) GetAllByUserId(id int) ([]port.Ticket, error) {
	tickets := []port.Ticket{}
	// err := c.db.Find(&tickets).Error
	err := c.db.Where("user_id = ?", id).Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (c ticketRepo) GetById(id int) (*port.Ticket, error) {
	ticket := port.Ticket{}
	err := c.db.First(&ticket, id).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (c ticketRepo) Create(ticket port.Ticket) (*port.Ticket, error) {
	err := c.db.Create(&ticket).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (c ticketRepo) Update(id int, ticket port.Ticket) error {
	err := c.db.Model(&port.Ticket{}).Where("ticket_id = ?", id).Updates(ticket).Error
	if err != nil {
		return err
	}
	return nil
}

func (c ticketRepo) Delete(id int) error {
	err := c.db.Delete(&port.Ticket{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c ticketRepo) Search(ticketName string) ([]port.Ticket, error) {
	ticket := []port.Ticket{}
	result := c.db.Find(&ticket, "ticket_name LIKE ?", "%"+ticketName+"%")
	if result.Error != nil {
		return ticket, errors.New("ticket not found")
	}
	return ticket, nil
}

func (c ticketRepo) UpdateStatusTicket(id int, status bool) error {
	err := c.db.Model(&port.Ticket{}).Where("ticket_id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}
