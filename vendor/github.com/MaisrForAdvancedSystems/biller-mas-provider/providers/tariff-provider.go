package providers

import (
	"context"
	"errors"
	"fmt"
	"log"

	// Proto Service Tariff

	"github.com/HassanAbdelzaher/lama"
	"github.com/MaisrForAdvancedSystems/biller-mas-provider/dbcontext"
	tool "github.com/MaisrForAdvancedSystems/biller-mas-provider/tools"
	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
	tarifserve "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
	dbmodels "github.com/MaisrForAdvancedSystems/mas-db-models/dbmodels"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TariffProvider struct..
type TariffProvider struct {
	InfoProvider
}

func loadCTG() ([]*tarifserve.Ctg, error) {
	var db *lama.Lama = dbcontext.DbConnPool
	query := fmt.Sprintf(`select * from dbo.CTG_CONSUMPTIONTYPES`)
	var ctgdata []*dbmodels.CTG_CONSUMPTIONTYPES
	err := db.DB.Unsafe().Select(&ctgdata, query)
	if err != nil {
		return nil, err
	}
	ctgs := make([]*tarifserve.Ctg, 0)
	if ctgdata == nil {
		return nil, errors.New("Not Found CTG")
	}
	for ids := range ctgdata {
		ctgelment := ctgdata[ids]
		if ctgelment != nil {
			var cgroup dbmodels.CTG_CONSUMPTIONTYPEGRPS
			log.Println("load ctg:", ctgelment.CTYPE_ID)
			err = db.Where(dbmodels.CTG_CONSUMPTIONTYPEGRPS{
				CTYPEGRP_ID: ctgelment.CTYPEGRP_ID,
			}).First(&cgroup)
			if err != nil {
				return nil, err
			}
			groupDescription := ""
			if cgroup.DESCRIPTION != nil {
				groupDescription = *cgroup.DESCRIPTION
			}
			// tarif services
			ctgstarifs := make([]*tarifserve.ServiceTariff, 0)
			query = fmt.Sprintf(`select * from dbo.CTG_CONSUMPTION_SERVICETARIFF where CTYPE_ID ='%s' order by SERVICE_TYPE`, ctgelment.CTYPE_ID)
			var ctgtarifsdata []*dbmodels.CtgConsumptionServicetariff
			err := db.DB.Unsafe().Select(&ctgtarifsdata, query)
			if err != nil {
				return nil, err
			}
			for idst := range ctgtarifsdata {
				ctgtarifelment := ctgtarifsdata[idst]
				if ctgtarifelment != nil {
					var ServiceTypectg tarifserve.SERVICE_TYPE = tarifserve.SERVICE_TYPE(ctgtarifelment.ServiceType)
					ctgtarif := tarifserve.ServiceTariff{ServiceType: &ServiceTypectg, TariffCode: &ctgtarifelment.TariffID, TransCode: &ctgtarifelment.TransCode, DiscountPercentage: ctgtarifelment.DiscountPercentage, IsZeroTarif: ctgtarifelment.IsZeroTarif, TaxPercentage: ctgtarifelment.TaxPercentage}
					ctgstarifs = append(ctgstarifs, &ctgtarif)
				}
			}
			ctg := tarifserve.Ctg{CType: &ctgelment.CTYPE_ID, CTypeGroupid: &ctgelment.CTYPEGRP_ID, Tariffs: ctgstarifs, Description: ctgelment.DESCRIPTION, GroupDescription: &groupDescription, Weigth: ctgelment.WEIGHT}
			ctgs = append(ctgs, &ctg)
		}
	}
	return ctgs, nil
}

func loadTariffs() ([]*tarifserve.Tariff, error) {
	var db *lama.Lama = dbcontext.DbConnPool
	var tarifdata []*dbmodels.Tariffs
	err := db.Model(dbmodels.Tariffs{}).Find(&tarifdata)
	if err != nil {
		return nil, err
	}
	tarifs := make([]*tarifserve.Tariff, 0)
	if tarifdata == nil {
		return nil, errors.New("Not Found")
	}
	for ids := range tarifdata {
		tarifelment := tarifdata[ids]
		if tarifelment != nil {
			// tarif bands
			tarifbands := make([]*tarifserve.TariffBand, 0)
			query := fmt.Sprintf(`select * from dbo.TARIFF_BANDS where TARIFF_ID = %s order by RANGE_FROM`, *tool.Int32ToString(&tarifelment.TarrifID))
			var tarifbandsdata []*dbmodels.TariffBands
			err = db.DB.Unsafe().Select(&tarifbandsdata, query)
			if err != nil {
				return nil, err
			}
			for idst := range tarifbandsdata {
				tarifbandelment := tarifbandsdata[idst]
				if tarifbandelment != nil {
					tarifband := tarifserve.TariffBand{From: &tarifbandelment.From, To: &tarifbandelment.To, Charge: &tarifbandelment.Charge, Constant: &tarifbandelment.Constant}
					tarifbands = append(tarifbands, &tarifband)
				}
			}
			// tarif extra fees
			tarifextras := make([]*tarifserve.ExtraTariffFess, 0)
			query = fmt.Sprintf(`select * from dbo.TARIFF_EXTRA_FESS where TARIFF_ID = %s `, *tool.Int32ToString(&tarifelment.TarrifID))
			var tarifextradata []*dbmodels.TariffExtraFess
			err = db.DB.Unsafe().Select(&tarifextradata, query)
			if err != nil {
				return nil, err
			}
			for idst := range tarifextradata {
				tarifextraelment := tarifextradata[idst]
				if tarifextraelment != nil {
					tarifextra := tarifserve.ExtraTariffFess{Description: tarifextraelment.Description, TransCode: &tarifextraelment.TransCode, FixedAmount: tarifextraelment.FixedAmount, AmountPerMeter: tarifextraelment.AmountPerMeter, TaxPercentage: tarifextraelment.TaxPercentage}
					tarifextras = append(tarifextras, &tarifextra)
				}
			}
			tarif := tarifserve.Tariff{TariffCode: &tarifelment.TariffCode, EffectDate: timestamppb.New(tarifelment.EffectDate), Bands: tarifbands, ExtraFees: tarifextras}
			tarifs = append(tarifs, &tarif)
		}
	}
	return tarifs, nil
}

func loadRegularChargess() ([]*tarifserve.RegularCharge, error) {
	var db *lama.Lama = dbcontext.DbConnPool
	query := fmt.Sprintf(`select * from dbo.REGULAR_CHARGES`)
	var regulardata []*dbmodels.RegularCharges
	err := db.DB.Unsafe().Select(&regulardata, query)
	if err != nil {
		return nil, err
	}
	regulars := make([]*tarifserve.RegularCharge, 0)
	if regulardata == nil {
		return nil, errors.New("Not Found")
	}
	for ids := range regulardata {
		regularelment := regulardata[ids]
		if regularelment != nil {
			// RelationChargeEntity
			relationchargedata := []*dbmodels.RegularRelationEntity{}
			query = fmt.Sprintf(`select * from dbo.REGULAR_RELATION_ENTITY where REGULAR_CHARGE_ID = %s AND RELEATION_TYPE = %v `, *tool.Int32ToString(&regularelment.RegularChargeID), int32(tarifserve.ReleationType_CHARGE_RELEATION))
			err = db.DB.Unsafe().Select(&relationchargedata, query)
			if err != nil {
				return nil, err
			}
			relationcharge := tarifserve.RegularChargeEntity{}
			if len(relationchargedata) > 0 {
				var chargeentitytype tarifserve.ENTITY_TYPE = tarifserve.ENTITY_TYPE(relationchargedata[0].EntityType)
				var chargevaluesdata []*dbmodels.RegularRelationValues
				query = fmt.Sprintf(`select * from dbo.REGULAR_RELATION_VALUES where REGULAR_CHARGE_ID = %s AND ENTITY_TYPE = %s `, *tool.Int32ToString(&regularelment.RegularChargeID), *tool.Int32ToString(&relationchargedata[0].EntityType))
				err = db.DB.Unsafe().Select(&chargevaluesdata, query)
				if err != nil {
					return nil, err
				}
				chargevalues := make([]*tarifserve.EntityChargeMappedValue, 0)

				for idst := range chargevaluesdata {
					chargevalelment := chargevaluesdata[idst]
					if chargevalelment != nil {
						chargevalue := tarifserve.EntityChargeMappedValue{From: chargevalelment.From, To: chargevalelment.To, LuKey: &chargevalelment.LuKey, Value: chargevalelment.Value}
						chargevalues = append(chargevalues, &chargevalue)
					}
				}
				relationcharge = tarifserve.RegularChargeEntity{EntityType: &chargeentitytype, MappedValues: chargevalues}

			}
			// RelationEnableEntity
			relationenabledata := []*dbmodels.RegularRelationEntity{}
			query = fmt.Sprintf(`select * from dbo.REGULAR_RELATION_ENTITY where REGULAR_CHARGE_ID = %s AND RELEATION_TYPE = %v `, *tool.Int32ToString(&regularelment.RegularChargeID), int32(tarifserve.ReleationType_ENABLE_RELEATION))
			err = db.DB.Unsafe().Select(&relationenabledata, query)
			if err != nil {
				return nil, err
			}
			relationenable := tarifserve.RegularEnableEntity{}
			if len(relationenabledata) > 0 {
				var enableentitytype tarifserve.ENTITY_TYPE = tarifserve.ENTITY_TYPE(relationenabledata[0].EntityType)
				var enablevaluesdata []*dbmodels.RegularRelationValues
				query = fmt.Sprintf(`select * from dbo.REGULAR_RELATION_VALUES where REGULAR_CHARGE_ID = %s AND ENTITY_TYPE = %s `, *tool.Int32ToString(&regularelment.RegularChargeID), *tool.Int32ToString(&relationenabledata[0].EntityType))
				err = db.DB.Unsafe().Select(&enablevaluesdata, query)
				if err != nil {
					return nil, err
				}
				enablevalues := make([]*tarifserve.EntityEnableMappedValue, 0)

				for idst := range enablevaluesdata {
					enablevalelment := enablevaluesdata[idst]
					if enablevalelment != nil {
						enablevalue := tarifserve.EntityEnableMappedValue{LuKey: &enablevalelment.LuKey, Value: enablevalelment.EnableValue}
						enablevalues = append(enablevalues, &enablevalue)
					}
				}
				relationenable = tarifserve.RegularEnableEntity{EntityType: &enableentitytype, MappedValues: enablevalues}
			}

			if len(relationenabledata) <= 0 && len(relationchargedata) <= 0 {
				return nil, errors.New("No Data For Enable-Charge REGULAR_RELATION_ENTITY")
			}

			var regularchargecalperiod tarifserve.RegularChargePeriod = tarifserve.RegularChargePeriod(regularelment.ChargeCalcPeriod)
			var regularchargetype tarifserve.ChargeType = tarifserve.ChargeType(regularelment.ChargeType)
			var regularservicetype tarifserve.SERVICE_TYPE = tarifserve.SERVICE_TYPE(regularelment.ServiceType)
			var effDateTo *timestamp.Timestamp = nil
			if regularelment.EffectDateTo != nil {
				effDateTo = timestamppb.New(*regularelment.EffectDateTo)

			}
			regular := tarifserve.RegularCharge{RegularChargeId: tool.Int32ToString(&regularelment.RegularChargeID), Bypass: regularelment.ByPass, ChargeInterval: regularelment.ChargeInterval,
				ChargeMonthlyDay: regularelment.ChargeMonthlyDay, FixedCharge: regularelment.FixedCharge, FixedChargeDiscount: regularelment.FixedChargeDiscount,
				IsChargable: &regularelment.IsChargable, MinCharge: regularelment.MinCharge, Title: &regularelment.Title, PerUnit: regularelment.PerUnit,
				TransCode: &regularelment.TransCode, VatPercentage: regularelment.VatPercentage, EffectDate: timestamppb.New(regularelment.EffectDate), EffectDateTo: effDateTo,
				ChargeCalcPeriod: &regularchargecalperiod, ChargeType: &regularchargetype, ServiceType: &regularservicetype}
			if len(relationenabledata) > 0 {
				regular.RelationEnableEntity = &relationenable
			}
			if len(relationchargedata) > 0 {
				regular.RelationChargeEntity = &relationcharge
			}
			if regularelment.CalcStartegy != nil {
				var regulatstrategy tarifserve.ChargeRegularCalcStrategy = tarifserve.ChargeRegularCalcStrategy(*regularelment.CalcStartegy)
				regular.CTypeCalcBase = &regulatstrategy
			}
			regulars = append(regulars, &regular)
		}
	}
	return regulars, nil
}

// GetSetupData implement Tarifs
func (s *TariffProvider) GetSetupData(cn context.Context, empty *tarifserve.Empty) (*tarifserve.SetupData, error) {
	ctgs, err := loadCTG()
	if err != nil {
		return nil, err
	}
	tarifs, err := loadTariffs()
	if err != nil {
		return nil, err
	}
	regulars, err := loadRegularChargess()
	if err != nil {
		return nil, err
	}
	res := tarifserve.SetupData{Tariffs: tarifs, Ctgs: ctgs, RegularCharges: regulars}
	res.TransCodes = make([]*tarifserve.TransCode, 0)
	for k, v := range dbmodels.FinancialTransCodes {
		var code string = string(k)
		var desc string = v
		res.TransCodes = append(res.TransCodes, &billing.TransCode{
			Code:        &code,
			Description: &desc,
		})
	}
	return &res, nil
}
