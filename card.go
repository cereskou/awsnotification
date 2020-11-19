package main

//image -
const (
	//32x32
	// SUCCEEDED   string = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAACXBIWXMAAA7DAAAOwwHHb6hkAAADjElEQVRYw7WXz08bVxDHP/NsBzBUocq6UhGHHNqQmARCqhRMDm16oVK5cHCr/hE9thJqjmmJeovEv9BK9BIJQaVegF7q9QFM09RgTr1VqheVSHGw8fpND5gI/17AzGm1ejuf786bN29GCGiTW86gicqsCJ8g3AWuA4OAAi+Bv1EyqqzbQ11JT3gHQfxKpwWJbGxEQsxj+ALoDai3iGVJKyyk4vncuQRMZpyoicpjMXwFhDmf+WpZtAX9Nn3Pex1YQCIbuyFhniHE6YYpWfWZS8Xzex0FJHZjH4jhV4RrdNOUfbXMpG7mN1sKqP75712HnxbhM306EubkYeq5E62G/XLgx797TcI8m9x2og0CpEe+69qetxcRN33yuGYLEtnYiER4cYFsbzTLslZwJcL3TU+Hz+3UrXzOAEiI+W7D/VeaTMXzC1rmmyYrwmKYB5DJjDMYGpB/zlBkOiXacuWVJtP3vKNTyf21RPihvlhVCvquMX0y21V4oRYOoBX9EeXfutW9pldmjQgPOzgtBYa/1mR6ohY+9dwZMj2yhvBOk4R8aBAm2jgt2JJ+qv7xfrWFFzWZvlsLn/zDGTK9soYw0vRACBPh6q3WCj7r3vE2gI3ETgwJs9AMbouaTI/VwbedoVBfa3jVrhvgalP40Rs4AKlb+ScNkVCWbUmTbnP4egc4wFXToliohMTWv64RoSzbI026d+rgmSrccCNI6phqM1FvAxLml0TW+aiFiKQtk3Rv18G3nKFQNDgceCnTe7GtlomoFGxZP3NHvd86efpw0xkKD5wJDkrGoGy3qdv9JiKrUy8aI3HapnOx4fBbsnEm+LGAbaPKWofLo99ckdWpP52Pm8L3YsMY1hHeP3PdUtaMLekKUOwookdW6iMxvRcbRlgH3jtH3SxqSVdMetw7wLIU4BrtN1dk9SQxLwgHy5I77h0YAOvrE8APIkIisprYjX15ITj4esw8bkjcUW9XLYsBP+6XED9dAA6WxdSot1vTEWlJH6HsXHpHpOzYkj5qaMncMa+gvs6h7F8ifF/LOueOeYUGAQCpuJdTnxm4FBH76jOTGvVy9aW4ttTG85ta1gdd3Q5lR8s8SMVrZ4KmAk4iUTnU+2p5Guh0tB/NnlYO9X6rGbHjcDr1l3PThGUew+dnHE5/tr4uuNVsP/d0fKq7edv01I3nUu0l9M14vq3Kmi3pSnrc+y+I3/8BXLibRUW3sMAAAAAASUVORK5CYII="
	// FAILED      string = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAACXBIWXMAAAuJAAALiQE3ycutAAADhElEQVRYw72Xz2+TdRzHX59nz57BoO1+4HAIKiNMDRFoR0xLSCQu42DiwbMXDoQtAfztDQ8mmohxk81ijRjl4MF/QRNGPJinXp7uoNGAh62js9ApiElhXbuPh311ndvTPpTV963tt9/36/vr831/hYAqDcQFlUPAIBAFHgc6zc9/AtPAFHAFJe1k3KUg/Upd42gijDACDAN9AXlngYvABcdz/2gIoBSNCyLHgfeBHhrTLeBtVU21ZdJLgQFKsUQE+BJ4kY3Rt8BLjufO1wUoxRLbgW+AA2ysrgJDjudmfQHMyL9rgnk1xBHHcwv/fGFVbTYx094sc4B+4OuFaKJlDQDC8Q1c81p6ToSXVy2BOWrXfHe7CLJrJ5qdDWbR1QnlMtz5y6/FHWCv47k3LYMxUsu85a3XsL/6Anl6XwDzLuxPP8ZOnodwyK9VGHgVwCoNxMUUmfX9d+3EeuF52LIFe2IMefKJOuYTyJ4+pGcbEonUQj1RiiVaLVNefSucZmcpv/ImFIsQDmFf+AjZ01fTnMI85eEz6Oz1WgAPAUctU9trSr2pFYiODuzUOPLYozXMT6Mz2SC7ZdAyFwvBIe5CdzctqQnkkR0PYg4QlVIs8QPwTNB/SOwg9vgotG9G536De/eQvt2NmANcs6quVILPxBtQvIvs6H0Qc4AOq5FKoj/+hP78y8rn8iKoNtKVWCZMBNemNuzRc8hAFArzaG4O6e3F/jyF9O+9X4DblkkywdTejj3+IXI4/u+0V06eQnNz0N2F/VkSObj/fgBmLBOj6mvrVuzkGHJoYNWaa/4GlZOn0es5CIWwk+eXAYNpygKu1G0WiWB/Mo4c2L/uhtN8nsrwmWWIzZuwx85hHRsMAjBpoaRNhltfnZ3L53zfUzV3+yqI1lZa3n0Ha6gmxC1VvWyZ9HrRt1mlDItlKBTqHjXN56mMLEPo9DRLXqYWwKW2THpBTBLqAn71rQnhEBKJ1KvtK2fr4e3o4iL87huIi0C/47k5qYpjp4Ak/4/OOp773qpEpKopk16bLVdVP/ALpduA7012a4aywGHHc3NrMyFgcvuQSa/NMB+qNl8DYCCywBFgciOn3Yx8zcDWvYwczy2ocgx43QTIRlUEzir67H9HHvxxGkv0mAB5wsSooG/CS8Con3FggBWQeCvI0arn+W6gw/RxG5gx98qkql5uy6QXgvT7N8MQa0AqNNLSAAAAAElFTkSuQmCC"
	// WARNING     string = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAACXBIWXMAAAsTAAALEwEAmpwYAAADo0lEQVRYw8VXX2hTZxQ/35/c/DOtbUITxXktSDVVSqf9gxKJ6IN0VFZni0/qw0p1OPogsji2KVbE1gdFQXRdFXSwwiz4Ui0+WCwMxZjaIprWTjHpRJrZmDXNcpube7/PFw1Ka3MTQ/N7PPf78zvnfOfc30GgEZNHGovlEW89j0+7galfAmciYFIMAABMfQMIBwGTIWSyDAjOml5bW88bLeeidAtCre616vjIYRafbgIlKWhiS3UyNlquEdHZbj838DgrAuFTzZbE/Zsn2VR4PzCVQDbARMWF1ov62q9+tP7QNa2ZQKh5Xbny6tl1PhMvgxwAGUxjdOnKHfauh/60BEItVRuVoL+XJ+UiyCGQTohQsbze3um7+0kCoeZ15cl/Rv+CdJcTOooM5lsfmvjM/9tAVVbPu08nROgXq1yOriH/LALhjm8LZgauPdAY9t+X9St7PjS83EKvAsBuLekwuJuqrZ5LUQAA/P5DwtvXnquczwc+Ey9LePvaU+/0XalVsKlwCywQ2FS4JdTqrkgRUIJ+T9allhUDlShBvwcAAE/+vMPKpVgjLDC4FGt8/VODFctPB7dr7nC5hJIUkk8Ht2Mej7ohT+DStBsDY5X5IgBMrcTAuZg3ApyLGBAqyhsBhIowIAT5AwIMTA3nLwUsjAHhQB5TEKBA6BAArM9QaNhe7rTXfGSbem0DxjIjQHRDFJks/Twebc6wfOogEqr7fKFi7seCs/YGooK04OGnOomuqrqBbcf+jCKTpTtj9qaCHmIXq4mjtBqZC3sylovGRd0lJ65HMQAALV3bAYTK2i+3DBo2Nexa0v3ct+SPv30G19e7kMkyqD33RKZieUfqd1xy5vYYLrSd1V49rNfquZx6cVbPZcZVtVez9wW2syXnBsY+UkSGzU1HkcE8rCkCmKyZZSN0jcaHN6x3NRydU5RO7KtZoQaf3ONywjG/C5gRR+mBgj2/dAIARK8eb1EnXpwHxvD8ylg/QUTnBkenL/BJWT6xr7ZCDTzu48nE0rTeCMZJAAAuS7b0slz/iix31jl+8z1KP5h8t0FUxkd6uBSryknDMy7y0WVljfZfvcFZwZxrg/3CvaDpm1YXLna0AaHZ9whCJVzsaDM1HHDNdbmm4fTfw/XLldEHB7kU28uTicXapiD9f8hovkLL1p8uOdU3/lnTcWo8P/29Xvbe3Mpjkc2gqpXA2QrApOhda44AwgHAZBhZFt8Rqrbdth26mNBy7lu4R1q2mwyMIgAAAABJRU5ErkJggg=="
	//24x24
	// SUCCEEDED string = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAACXBIWXMAAA7DAAAOwwHHb6hkAAACTElEQVRIx7WWPU8UURSGn3NniREsuYmtEoVd4mojZLZzqbVRIyDyA+gNEnsjrYW0shSsERuoWTsnmGiC0RE2AUuNQ0UizTJzLHbZL2eHFeF0d+7Mc+49854PIcFc3w6IwzhCHsgAtrYVAF9R3mlI0csEO50Y0gF8SVLMI9wFDMkWobzVkCdeOtg91oG7bR+JYQHo49/st0bMeINBoflhy+ly23ZWDIUTwAH6xLCY27azsTdwt+y0OCxyCqYR095gsFR34H6zlyXF5xOePD5cFa55meC7ARCH5yeGK1+0wjjKQUu4UswDiOvbAemh3IVa4uGHjHmZ4Jfr27ykWEPoPVKXVrhixOHB/8IBvExQ0ogXzQISh3FTS6I4de8kwsMGvPYfJ8ThcVsS5A0wHKOC0uG+ZjXkaQy+Ck+3wCclxRLgtL07bJrSvw4P9/X2h5t7B95Q8KzNSSd4IQYOYP+KvQhqzjXWTU6q8KEmuJ8Ir/JyZfsD4WJbjNfDA72zcWOvLr3Rzf7ejeuNtevbSelJhqP8NIAfUwLHnPOyOvKp/0hytMC37MNj4VXzDUopdssw5vTJ2sjHhhOAXNlO1UqK04WUS0ZDikAUGz9DPnVBVkc3q05yZTuF8KorOEQaUpTah28Q7iWcZB14jbDQJRyUlfdXg/tnV+wOyXrpYNcAeOlgVyNmTgmOhswcdbd6HniDQUEj5k6hF8x5Q42uFtcyp8Xw8kxaZv0mFbIoK53UFdP0V7RCth3ecapoG1smEG4Bw0itbmnL2LKcNLb8AZySCLbXZydaAAAAAElFTkSuQmCC"
	// FAILED    string = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAACXBIWXMAAAuJAAALiQE3ycutAAACjUlEQVRIx62WT0gUYRjGf9+34yj0xyA0AvVeeogxg1HpYEkdyhAPQdGCkrBUGBhEJAopFgqWCeEhvATdhJYgWNza/hwaOjRdtHsnCQsvmjrszHTYb3PXnVlX7bntfDO/d75n3+95R1BEjmEeBC4BZ4AGoAoQwBLwHUgBcd22lsMYIgRcCQwAMeAAxfUHeA6M6Lb1e9sCjmGeBl4CNexMP4GobltzuRflFngXkNwFHOAI8MYxzGuBO1BvngR09qY0cFG3rcS/Asrz+V2+eZCWgHrdtpayFg0EwcVJA8rLwzFCIFrMoJUq4AGAVK0YK4Q3oj2bRBsfhbKyQHikvw9tagLZ3hZUpNsxzGqp+rygFf35efyv3xCtzUQeDYOmFcDllct4iTm81MegAhVAV2TwaG0/cKLwr3Lx3n1ANhxHtjYj6mrw338CyIO7QyPgumEmrmvqhIYsr5Puv4v2eBx5rh02HFhZKRUO0CAcw/wB1BXtiYoKtCdjiFNNAHiJJO7Q8HZwgGUZFhdbd+J9/rL5202X8hQAmurZ2mI3yehVIn038BJJWFtDdnbAvv249wcztoXrl6ZS0QiF93YTifVueu57meudHYinE6Tv3IPV1bDHF6SK3EBFbsXy4a4Lno87Oob36jWiqRFtegoOVYYhUhKIq8jNP0ctJrI7Gtwtvo/7cDxTpP4YsuNCYOoDs9ksmgRuF9jT3pY5RGHdIgTybBve2xT4/tbVGd22rmcLHAYWVOT+Dy2rsFuUAGoSRVXU7lUe0KPb1mLewFGTqGePRTzgpm5b8WIj8zzwQkXuTm3pyYUXjEy1kwRQD0wD6yWAHWBGeR4v6asiZzfVQFfOZ0t19oSqpkgBs1m/g/QXn1DubraXMmgAAAAASUVORK5CYII="
	// WARNING   string = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAACXBIWXMAAAsTAAALEwEAmpwYAAACkklEQVRIx7WWv08TYRjHv+9z10JLjvJDynGt1BggCoNVJ3UB4wwDDbK4UBhMHEwksvEHkJg4mBiDbUhMDCbtIC5MMGmMCQYGiBETpCnNXaMlUKD0x93rZBXt3QGh3+nN5bnP8+Oe3PdlsJA6dr1TV7/fRbHQx7nezcC8AMDB04wJa3A4FwX5wht5emndjMEqgsPBDj2dmOIH2QFwTlZFgDGDuaW3grf9sRxZ/mabIBXyjxg7P55BL7lwEglijjznHiixZPTvx3QU7ps0MmrkxHAA0EsuI6NGUiHfZMUOUiH/iJHRIgCvPAl3/T1qal0GACOjBfnB7iuzqVOjN6zEt6LlBGo42FlKfFmxqfyWf6H0AQCSt8WbAN5bjUtsv3RFjiyvEwDoWmLqVGOxGJeuJaYAgNTRa108l+3HGYvnsv1q+GoX6drmsO0qnioDJz2dGCYU872olor5XuLc6K4Wn3OjmwDWUrUOwFoIZot/Vh8C4OnqJUBaZCSscqDNNtRZG0gOerYBAPs7ARQO7QfEhFURomMB+dwd2+jC4evjQI9IFBdIVDpmwUi3L4cM5qx9x5y1c2BkHKN8XWi7OEutLz5tMLcUs4snqWHcN7/X75vfGyCpcdyW75Ji8vTSBgGAIAcmIDiylm/UNcyVz25pzsYbsoL3/ETZD+Tpz5vkaR4DY+at57JD5fPh/pCVw1F985gcXdn8z9FSIf9DY1t7UvHfxBhnNa5FAOD5XB84ZxXhDd5HSnzrqZVlhozdny9RKnpOtjGOHZKaRpX4VszUMgFAiSVjYqCnh9V5ZkBUsAUTFVhd/YzYfrnnX7jpreK3tPs3lFLy6yCK+T7OjT/XFs7TjGgNjppF0d8Zb33+MWXG+AXRn/mwCw/nuAAAAABJRU5ErkJggg=="
	//20x20
	SUCCEEDED string = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAAUCAYAAACNiR0NAAAACXBIWXMAAA7DAAAOwwHHb6hkAAACAElEQVQ4y6WVvWtTYRSHn/Pmw9j5BpqoU4PXVAWHQJNOTRFcOwWMu1QqTvpX6ChocDaFTnF1SN1uhQxCP8SCSiO0hZtFh9qP5B6H3LTXaxLT9Le9h/c8nHP4nfcV+ii/biUkLiURFoAcQhpQlD2goUpNT3Rl7VbrMJwr4UDhS/K+GF4AaYZrVz2eOra73Bc488mKRCbkJcIi55HyunOgjz/eaXUAor34WLBuSYuRCQF4dFqh32aVC0g9yo7tLkt+3UqYhHwdYWbBNpva4YlEeIOQ7M3UO9IpI3EpnQsGTfUoOln3nba5i+L68bSJSSnqW2N0WIc554b73T8fAEeBeS4YIBdqZ1+P9R7KtxBsR70zWGHLykiUVYSrgTs5g5AKwrwTnXdutt5rW4sB6I56FB27C8tvWhmJShgGQsqEbS4RUQBnutXUts6hrBKEbVgZE+sD6yFmt5NNhGt/tdxm3pl2P/+zkhtWxsQHw1B+GKARqnJSotTzm1Y2GJ7dTl43l+TDQFhXDaNKrY/7J01M6oWtZLYHQ6gDV/7jz9pwYyv7KA8xvBoBtusd6lRv9cpieHvB1Xvg2G7VADi2W0WpjE+j4thuFeDUNu2fujQWVKm0f+nSsAe2LMJz/5UeOjNVnvUqGwj0v4DLJi4lzr6AlA/ZAxooNe9YV9Zut36Hc/8Aq3DRkJ+SQMAAAAAASUVORK5CYII="
	FAILED    string = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAAUCAYAAACNiR0NAAAACXBIWXMAAAuJAAALiQE3ycutAAACC0lEQVQ4y6WVT2gTQRSHv5lZ9tKm9xR6LeJB6ELBvSRVK0ugKL0XehAUVDBCsCDWmlitUlBIrVpPgjdBlCCCRbyuYLoKSk/aa6/+odCNTtZDt012s9m0+I5ved/8Zt7vvRUkhD9iCyEYByaBUWAQEMAmUAdeAqum5zbjtSKeaFh2DlgCjpAe68Al03PfdQU2LHsGuA1I9hcBUAHKpucGAEYMdoeDhQDmQgHX9xQ2LDsPvD+AsiSlBdNz30p/xBZANQJTCllwulbL42OQ6Y8rrfqWrWTYzUgDpHMSNT+HunCuEzZ5CnV3HpnPxT8NC5hQs9mhUmiNlv5v3xF9fcjpKdCa4NPnFuzqFfTyCs3nLxIdZ8RhO8QAfb+6c/vzZ8H3YWurBXv6rNtrjBqhaROfWd8LocWLAOgHj9NgAIMyydwRbn0NgqCtmek2MsJxyiZ288Qx1K0b6OUV2N5GlYqIgQH00qO2QyKxaYSzaXXACg6qfA398EnkmqpUhEwGvbAIzY5RrqvZ7NBfYCoCG8uhFiodsODrOvz6jTozjVCK4ONaHFgRDeuoBPEFOLyXzvQj8zmar98kG7vg0HQ/wI+f7ekN4NDu6I0Dq70a1CNOm55bkwDhCqr8B2zR9NwasWVQBm7uxxtxWEAwk7ZgnXBZDPcAbQCXd5V1BQL4lq0ETKT8Al4BNdNz/8Rr/wFXIamunVPe6gAAAABJRU5ErkJggg=="
	WARNING   string = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAAUCAYAAACNiR0NAAAACXBIWXMAAAsTAAALEwEAmpwYAAACMUlEQVQ4y61VS2wSURQ97zEDDGI/qBQ61miiuDIRa9waE1s3VKws6sZu2l2jDdS4cd8E/ATdVjd1ZZqGYJD42dRFXdWoCxOiqTEpjEWTAsUCwzDvuZD64TM1kbN6ue/ec++7ue9cghb4eu2cVPv4ZoyrJT9nbBCAu371hVD6mlhsceGI95Hz1rNyYyxpNCgB+TIr5iKoVV0wgmBep7t7r/cvZh7+aabbh3x0Ssj4HHMsl53fkQwAalUXy2XnMz7HXD46JTRVmPE57vPS5kRToMn0kojW5wDAtcowdP100zNtXQ/kxMbkrwqVgDzekgwAdH1ZThZm5WRhFrq+3MqFlzYnlIA8DgA0GxqSWDEXxn+CFXORbGjIRvXVd5f+qWc797RPX307Rrla8qND4GrZT+tz1hlCxgYpwN3oGLibAuDoICMFiNI5PqIIhJpWOHCgvQ8dSQ9L++sz6QVnBq50RSBWKcar5YsGnT4GRtwgADjba1igRYpRwXNygYjmdNusomXJ7D0zIJ44O0BEy1JbP8GcFg8fX6DOSFIl9t6ZtlkpTTpvPq30hZ9UQEiyLaG9Z8Z5+4X6WxxG9tzlW4WrLSr8QF2HzgMAW//0mGtVTwtxuCcnNqb/ki8pEAySXd3RRonkmurR11IpfS2VaiYjILauqHThStBIYEf59/wdrqkHDQdENH8m9p5Q/6ISM1RsAPh2Y1TU3r/y88pWfQVs/ybycwVYbXHx6Kn4vnBCa4z9Ab8/3TDwZ13gAAAAAElFTkSuQmCC"
)

//Fact -
type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//Input -
type Input struct {
	Type        string `json:"@type"`
	ID          string `json:"id"`
	IsMultiline bool   `json:"isMultiline"`
	Title       string `json:"title"`
}

//Action -
type Action struct {
	Type   string `json:"@type"`
	Name   string `json:"name"`
	Target string `json:"target"`
}

//PotentialAction -
type PotentialAction struct {
	Type    string               `json:"@type"`
	Name    string               `json:"name"`
	Inputs  []*Input             `json:"inputs,omitempty"`
	Actions []*Action            `json:"actions,omitempty"`
	Target  string               `json:"target,omitempty"`
	Targets []*map[string]string `json:"targets,omitempty"`
}

//Section -
type Section struct {
	ActivityTitle    string  `json:"activityTitle"`
	ActivitySubtitle string  `json:"activitySubtitle"`
	ActivityImage    string  `json:"activityImage"`
	Facts            []*Fact `json:"facts"`
	Text             string  `json:"text"`
}

//Card -
type Card struct {
	Type             string             `json:"@type"`
	Context          string             `json:"@context"`
	Summary          string             `json:"summary"`
	ThemeColor       string             `json:"themeColor"`
	Title            string             `json:"title"`
	Text             string             `json:"text"`
	Sections         []*Section         `json:"sections"`
	PotentialActions []*PotentialAction `json:"potentialAction"`
}

//NewCard -
func NewCard() *Card {
	return &Card{
		Type:             "MessageCard",
		Context:          "https://schema.org/extensions",
		Sections:         make([]*Section, 0),
		PotentialActions: make([]*PotentialAction, 0),
	}
}

//NewAction -
func (c *Card) NewAction(name, actiontype string) *PotentialAction {
	action := &PotentialAction{
		Name:    name,
		Type:    actiontype,
		Targets: make([]*map[string]string, 0),
	}
	c.PotentialActions = append(c.PotentialActions, action)

	return action
}

//AddTarget -
func (a *PotentialAction) AddTarget(target map[string]string) {
	a.Targets = append(a.Targets, &target)
}

//NewSection -
func (c *Card) NewSection() *Section {
	section := &Section{
		Facts: make([]*Fact, 0),
	}
	c.Sections = append(c.Sections, section)

	return section
}

//NewFact -
func (s *Section) NewFact(name, value string) *Fact {
	fact := &Fact{
		Name:  name,
		Value: value,
	}
	s.Facts = append(s.Facts, fact)

	return fact
}
